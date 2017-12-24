package ts29338

import (
	"bytes"
	"encoding/binary"

	dia "github.com/fkgi/diameter"
	"github.com/fkgi/sms"
	"github.com/fkgi/teldata"
)

func setUserName(v teldata.IMSI) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 1, VenID: 0, FlgV: false, FlgM: true, FlgP: false}
	a.Encode(v.String())
	return
}

func getUserName(a dia.RawAVP) (v teldata.IMSI, e error) {
	s := new(string)
	if a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		v, e = teldata.ParseIMSI(*s)
	}
	return
}

func setMSISDN(v teldata.E164) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 701, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	a.Encode(v.Bytes())
	return
}

func getMSISDN(a dia.RawAVP) (v teldata.E164, e error) {
	s := new([]byte)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		v, e = teldata.B2E164(*s)
	}
	return
}

// MTType indicate SM-RP-MTI
// SM-RP-MTI AVP contain the RP-Message Type Indicator of the Short Message.
type MTType int

const (
	// UnknownMT is no SM-RP-MTI
	UnknownMT MTType = iota
	// DeliverMT is SM_DELIVER
	DeliverMT
	// StatusReportMT is SM_STATUS_REPORT
	StatusReportMT
)

func setSMRPMTI(v MTType) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3308, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case DeliverMT:
		a.Encode(dia.Enumerated(0))
	case StatusReportMT:
		a.Encode(dia.Enumerated(1))
	}
	return
}

func getSMRPMTI(a dia.RawAVP) (v MTType, e error) {
	s := new(dia.Enumerated)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e != nil {
	} else if *s == 0 {
		v = DeliverMT
	} else if *s == 1 {
		v = StatusReportMT
	} else {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
	}
	return
}

// SM-RP-SMEA AVP contain the RP-Originating SME-address of
// the Short Message Entity that has originated the SM.
// It shall be formatted according to the formatting rules of
// the address fields described in 3GPP TS 23.040.
func setSMRPSMEA(v sms.Address) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3309, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	_, b := v.Encode()
	a.Encode(b)
	return
}

func getSMRPSMEA(a dia.RawAVP) (v sms.Address, e error) {
	s := new([]byte)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		v.Decode(byte(len(*s)-3)*2, *s)
	}
	return
}

// SRR-Flags AVP contain a bit mask.
// GPRS-Indicator shall be ture if the SMS-GMSC supports receiving
// of two serving nodes addresses from the HSS.
// SM-RP-PRI shall be true if the delivery of the short message shall
// be attempted when a service centre address is already contained
// in the Message Waiting Data file.
// Single-Attempt if true indicates that only one delivery attempt
// shall be performed for this particular SM.
func setSRRFlags(g, p, s bool) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3310, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	i := uint32(0)
	if g {
		i = i | 0x00000001
	}
	if p {
		i = i | 0x00000002
	}
	if s {
		i = i | 0x00000004
	}
	a.Encode(i)
	return
}

func getSRRFlags(a dia.RawAVP) (g, p, s bool, e error) {
	v := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(v); e == nil {
		g = (*v)&0x00000001 == 0x00000001
		p = (*v)&0x00000002 == 0x00000002
		s = (*v)&0x00000004 == 0x00000004
	}
	return
}

// RequiredInfo indicate SM-Delivery-Not-Intended AVP data
// SM-Delivery-Not-Intended AVP indicate by its presence
// that delivery of a short message is not intended.
type RequiredInfo int

const (
	// LocationRequested is no SM-Delivery-Not-Intended
	LocationRequested RequiredInfo = iota
	// OnlyImsiRequested is ONLY_IMSI_REQUESTED
	OnlyImsiRequested
	// OnlyMccMncRequested is ONLY_MCC_MNC_REQUESTED
	OnlyMccMncRequested
)

func setSMDeliveryNotIntended(v RequiredInfo) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3311, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case OnlyImsiRequested:
		a.Encode(dia.Enumerated(0))
	case OnlyMccMncRequested:
		a.Encode(dia.Enumerated(1))
	}
	return
}

func getSMDeliveryNotIntended(a dia.RawAVP) (v RequiredInfo, e error) {
	s := new(dia.Enumerated)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e != nil {
	} else if *s == 0 {
		v = OnlyImsiRequested
	} else if *s == 1 {
		v = OnlyMccMncRequested
	} else {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
	}
	return
}

// Serving-Node AVP shall contain the information about
// the network node serving the targeted SMS user.

// NodeType is tyoe of serving node address
type NodeType int

const (
	// NodeSGSN present SGSN
	NodeSGSN NodeType = iota
	// NodeMME present MME
	NodeMME
	// NodeMSC present MSC
	NodeMSC
	// NodeIPSMGW present IP-SM-GW
	// NodeIPSMGW
)

func setServingNode(t NodeType, d teldata.E164, n, r dia.Identity) dia.RawAVP {
	a := dia.RawAVP{Code: 2401, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	return setSN(a, t, d, n, r)
}

func setAdditionalServingNode(t NodeType, d teldata.E164, n, r dia.Identity) dia.RawAVP {
	a := dia.RawAVP{Code: 2406, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	return setSN(a, t, d, n, r)
}

func setSN(a dia.RawAVP, t NodeType, d teldata.E164, n, r dia.Identity) dia.RawAVP {
	v := make([]dia.RawAVP, 1, 3)
	switch t {
	case NodeSGSN:
		v[0] = dia.RawAVP{Code: 1489, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
		v[0].Encode(d.Bytes())
		if len(n) != 0 {
			nv := dia.RawAVP{Code: 2409, VenID: 10415, FlgV: true, FlgM: false, FlgP: false}
			nv.Encode(n)
			v = append(v, nv)
		}
		if len(r) != 0 {
			rv := dia.RawAVP{Code: 2410, VenID: 10415, FlgV: true, FlgM: false, FlgP: false}
			rv.Encode(r)
			v = append(v, rv)
		}
	case NodeMME:
		v[0] = dia.RawAVP{Code: 1645, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
		v[0].Encode(d.Bytes())
		if len(n) != 0 {
			nv := dia.RawAVP{Code: 2402, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
			nv.Encode(n)
			v = append(v, nv)
		}
		if len(r) != 0 {
			rv := dia.RawAVP{Code: 2408, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
			rv.Encode(r)
			v = append(v, rv)
		}
	case NodeMSC:
		v[0] = dia.RawAVP{Code: 2403, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
		v[0].Encode(d.Bytes())
		if len(n) != 0 {
			nv := dia.RawAVP{Code: 2402, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
			nv.Encode(n)
			v = append(v, nv)
		}
		if len(r) != 0 {
			rv := dia.RawAVP{Code: 2408, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
			rv.Encode(r)
			v = append(v, rv)
		}
	}

	a.Encode(v)
	return a
}

func getServingNode(a dia.RawAVP) (t NodeType, d teldata.E164, n, r dia.Identity, e error) {
	return getSN(a)
}

func getAdditionalServingNode(a dia.RawAVP) (t NodeType, d teldata.E164, n, r dia.Identity, e error) {
	return getSN(a)
}

func getSN(a dia.RawAVP) (t NodeType, d teldata.E164, n, r dia.Identity, e error) {
	o := []dia.RawAVP{}
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else {
		e = a.Decode(&o)
	}
	for _, a := range o {
		b := new([]byte)
		switch a.Code {
		case 1489:
			if !a.FlgV || !a.FlgM || a.FlgP {
				e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
			} else if e = a.Decode(b); e == nil {
				d, e = teldata.B2E164(*b)
				t = NodeSGSN
			}
		case 1645:
			if !a.FlgV || !a.FlgM || a.FlgP {
				e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
			} else if e = a.Decode(b); e == nil {
				d, e = teldata.B2E164(*b)
				t = NodeMME
			}
		case 2403:
			if !a.FlgV || !a.FlgM || a.FlgP {
				e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
			} else if e = a.Decode(b); e == nil {
				d, e = teldata.B2E164(*b)
				t = NodeMSC
			}
		}
	}
	switch t {
	case NodeSGSN:
		for _, a := range o {
			switch a.Code {
			case 2409:
				if !a.FlgV || !a.FlgM || a.FlgP {
					e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
				} else {
					e = a.Decode(&n)
				}
			case 2410:
				if !a.FlgV || !a.FlgM || a.FlgP {
					e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
				} else {
					e = a.Decode(&r)
				}
			}
		}
	case NodeMME, NodeMSC:
		for _, a := range o {
			switch a.Code {
			case 2402:
				if !a.FlgV || !a.FlgM || a.FlgP {
					e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
				} else {
					e = a.Decode(&n)
				}
			case 2408:
				if !a.FlgV || !a.FlgM || a.FlgP {
					e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
				} else {
					e = a.Decode(&r)
				}
			}
		}
	}
	return
}

func setLMSI(v uint32) dia.RawAVP {
	a := dia.RawAVP{Code: 2400, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	a.Encode(buf.Bytes())
	return a
}

func getLMSI(a dia.RawAVP) (v uint32, e error) {
	s := new([]byte)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil && len(*s) == 4 {
		binary.Read(bytes.NewBuffer(*s), binary.BigEndian, &v)
	}
	return
}

func setUserIdentifier(v teldata.E164) dia.RawAVP {
	a := dia.RawAVP{Code: 3102, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	t := []dia.RawAVP{
		dia.RawAVP{Code: 701, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}}
	t[0].Encode(v.Bytes())
	a.Encode(t)
	return a
}

func getUserIdentifier(a dia.RawAVP) (v teldata.E164, e error) {
	o := []dia.RawAVP{}
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(&o); e == nil {
		for _, a := range o {
			if a.Code != 701 {
				continue
			}
			if !a.FlgV || !a.FlgM || a.FlgP {
				e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
			} else {
				a.Decode(&v)
			}
		}
	}
	return
}

// MWD-Status AVP contain a bit mask.
// SCAddrNotIncluded shall indicate the presence of the SC Address in the Message Waiting Data in the HSS.
// MNRF shall indicate that the MNRF flag is set in the HSS.
// MCEF shall indicate that the MCEF flag is set in the HSS.
// MNRG shall indicate that the MNRG flag is set in the HSS.
func setMWDStatus(sca, mnrf, mcef, mnrg bool) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3312, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}

	i := uint32(0)
	if sca {
		i = i | 0x00000001
	}
	if mnrf {
		i = i | 0x00000002
	}
	if mcef {
		i = i | 0x00000004
	}
	if mnrg {
		i = i | 0x00000008
	}
	a.Encode(i)
	return
}

func getMWDStatus(a dia.RawAVP) (sca, mnrf, mcef, mnrg bool, e error) {
	s := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		sca = (*s)&0x00000001 == 0x00000001
		mnrf = (*s)&0x00000002 == 0x00000002
		mcef = (*s)&0x00000004 == 0x00000004
		mnrg = (*s)&0x00000008 == 0x00000008
	}
	return
}

// AbsentDiag indicate Absent-User-Diagnostic-SM
type AbsentDiag int

const (
	// NoAbsentDiag is "no diag data"
	NoAbsentDiag AbsentDiag = iota
	// NoPagingRespMSC is "no paging response via the MSC"
	NoPagingRespMSC
	// IMSIDetached is "IMSI detached"
	IMSIDetached
	// RoamingRestrict is "roaming restriction"
	RoamingRestrict
	// DeregisteredNonGPRS is "deregistered in the HLR for non GPRS"
	DeregisteredNonGPRS
	// PurgedNonGPRS is "MS purged for non GPRS"
	PurgedNonGPRS
	// NoPagingRespSGSN is "no paging response via the SGSN"
	NoPagingRespSGSN
	// GPRSDetached is "GPRS detached"
	GPRSDetached
	// DeregisteredGPRS is "deregistered in the HLR for GPRS"
	DeregisteredGPRS
	// PurgedGPRS is "MS purged for GPRS"
	PurgedGPRS
	// UnidentifiedSubsMSC is "Unidentified subscriber via the MSC"
	UnidentifiedSubsMSC
	// UnidentifiedSubsSGSN is "Unidentified subscriber via the SGSN"
	UnidentifiedSubsSGSN
	// DeregisteredIMS is "deregistered in the HSS/HLR for IMS"
	DeregisteredIMS
	// NoRespIPSMGW is "no response via the IP-SM-GW"
	NoRespIPSMGW
	// TempUnavailable is "the MS is temporarily unavailable"
	TempUnavailable
)

// MMEAbsentUserDiagnosticSM AVP shall indicate the diagnostic
// explaining the absence of the user given by the MME.
func setMMEAbsentUserDiagnosticSM(v AbsentDiag) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3313, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case NoPagingRespMSC:
		a.Encode(uint32(0))
	case IMSIDetached:
		a.Encode(uint32(1))
	case RoamingRestrict:
		a.Encode(uint32(2))
	case DeregisteredNonGPRS:
		a.Encode(uint32(3))
	case PurgedNonGPRS:
		a.Encode(uint32(4))
	case UnidentifiedSubsMSC:
		a.Encode(uint32(9))
	case TempUnavailable:
		a.Encode(uint32(13))
	}
	return
}

func getMMEAbsentUserDiagnosticSM(a dia.RawAVP) (v AbsentDiag, e error) {
	s := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		switch *s {
		case 0:
			v = NoPagingRespMSC
		case 1:
			v = IMSIDetached
		case 2:
			v = RoamingRestrict
		case 3:
			v = DeregisteredNonGPRS
		case 4:
			v = PurgedNonGPRS
		case 9:
			v = UnidentifiedSubsMSC
		case 13:
			v = TempUnavailable
		default:
			e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
		}
	}
	return
}

// MSCAbsentUserDiagnosticSM AVP shall indicate the diagnostic
// explaining the absence of the user given by the MSC.
func setMSCAbsentUserDiagnosticSM(v AbsentDiag) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3314, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case NoPagingRespMSC:
		a.Encode(uint32(0))
	case IMSIDetached:
		a.Encode(uint32(1))
	case RoamingRestrict:
		a.Encode(uint32(2))
	case DeregisteredNonGPRS:
		a.Encode(uint32(3))
	case PurgedNonGPRS:
		a.Encode(uint32(4))
	case UnidentifiedSubsMSC:
		a.Encode(uint32(9))
	case TempUnavailable:
		a.Encode(uint32(13))
	}
	return
}

func getMSCAbsentUserDiagnosticSM(a dia.RawAVP) (v AbsentDiag, e error) {
	s := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		switch *s {
		case 0:
			v = NoPagingRespMSC
		case 1:
			v = IMSIDetached
		case 2:
			v = RoamingRestrict
		case 3:
			v = DeregisteredNonGPRS
		case 4:
			v = PurgedNonGPRS
		case 9:
			v = UnidentifiedSubsMSC
		case 13:
			v = TempUnavailable
		default:
			e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
		}
	}
	return
}

// SGSNAbsentUserDiagnosticSM AVP shall indicate the diagnostic
// explaining the absence of the user given by the SGSN.
func setSGSNAbsentUserDiagnosticSM(v AbsentDiag) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3315, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case IMSIDetached:
		a.Encode(uint32(1))
	case RoamingRestrict:
		a.Encode(uint32(2))
	case NoPagingRespSGSN:
		a.Encode(uint32(5))
	case GPRSDetached:
		a.Encode(uint32(6))
	case DeregisteredGPRS:
		a.Encode(uint32(7))
	case PurgedGPRS:
		a.Encode(uint32(8))
	case UnidentifiedSubsSGSN:
		a.Encode(uint32(10))
	case TempUnavailable:
		a.Encode(uint32(13))
	}
	return
}

func getSGSNAbsentUserDiagnosticSM(a dia.RawAVP) (v AbsentDiag, e error) {
	s := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		switch *s {
		case 1:
			v = IMSIDetached
		case 2:
			v = RoamingRestrict
		case 5:
			v = NoPagingRespSGSN
		case 6:
			v = GPRSDetached
		case 7:
			v = DeregisteredGPRS
		case 8:
			v = PurgedGPRS
		case 10:
			v = UnidentifiedSubsSGSN
		case 13:
			v = TempUnavailable
		default:
			e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
		}
	}
	return
}

// AbsentUserDiagnosticSM AVP shall indicate the diagnostic explaining the absence of the subscriber.
func setAbsentUserDiagnosticSM(v AbsentDiag) (a dia.RawAVP) {
	a = dia.RawAVP{Code: 3322, VenID: 10415, FlgV: true, FlgM: true, FlgP: false}
	switch v {
	case NoPagingRespMSC:
		a.Encode(uint32(0))
	case IMSIDetached:
		a.Encode(uint32(1))
	case RoamingRestrict:
		a.Encode(uint32(2))
	case DeregisteredNonGPRS:
		a.Encode(uint32(3))
	case PurgedNonGPRS:
		a.Encode(uint32(4))
	case NoPagingRespSGSN:
		a.Encode(uint32(5))
	case GPRSDetached:
		a.Encode(uint32(6))
	case DeregisteredGPRS:
		a.Encode(uint32(7))
	case PurgedGPRS:
		a.Encode(uint32(8))
	case UnidentifiedSubsMSC:
		a.Encode(uint32(9))
	case UnidentifiedSubsSGSN:
		a.Encode(uint32(10))
	case DeregisteredIMS:
		a.Encode(uint32(11))
	case NoRespIPSMGW:
		a.Encode(uint32(12))
	case TempUnavailable:
		a.Encode(uint32(13))
	}
	return
}

func getAbsentUserDiagnosticSM(a dia.RawAVP) (v AbsentDiag, e error) {
	s := new(uint32)
	if !a.FlgV || !a.FlgM || a.FlgP {
		e = dia.InvalidAVP(dia.DiameterInvalidAvpBits)
	} else if e = a.Decode(s); e == nil {
		switch *s {
		case 0:
			v = NoPagingRespMSC
		case 1:
			v = IMSIDetached
		case 2:
			v = RoamingRestrict
		case 3:
			v = DeregisteredNonGPRS
		case 4:
			v = PurgedNonGPRS
		case 5:
			v = NoPagingRespSGSN
		case 6:
			v = GPRSDetached
		case 7:
			v = DeregisteredGPRS
		case 8:
			v = PurgedGPRS
		case 9:
			v = UnidentifiedSubsMSC
		case 10:
			v = UnidentifiedSubsSGSN
		case 11:
			v = DeregisteredIMS
		case 12:
			v = NoRespIPSMGW
		case 13:
			v = TempUnavailable
		default:
			e = dia.InvalidAVP(dia.DiameterInvalidAvpValue)
		}
	}
	return
}

/*
// SMDeliveryOutcome AVP contains the result of the SM delivery.
type SMDeliveryOutcome struct {
	E dia.Enumerated
	I uint32
}

// ToRaw return AVP struct of this value
func (v *SMDeliveryOutcome) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3316, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	// a.Encode()
	return a
}

// MMESMDeliveryOutcome AVP shall indicate the outcome of
// the SM delivery for setting the message waiting data
// in the HSS when the SM delivery is with an MME.
type MMESMDeliveryOutcome struct {
	SMDeliveryCause        SMDeliveryCause
	AbsentUserDiagnosticSM AbsentUserDiagnosticSM
}

// ToRaw return AVP struct of this value
func (v *MMESMDeliveryOutcome) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3317, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	if v != nil {
		t := []dia.RawAVP{
			v.SMDeliveryCause.ToRaw(),
			v.AbsentUserDiagnosticSM.ToRaw()}
		a.Encode(dia.GroupedAVP(t))
	}
	return a
}

// MSCSMDeliveryOutcome AVP shall indicate the outcome of
// the SM delivery for setting the message waiting data
// in the HSS when the SM delivery is with an MSC.
type MSCSMDeliveryOutcome struct {
	SMDeliveryCause        SMDeliveryCause
	AbsentUserDiagnosticSM AbsentUserDiagnosticSM
}

// ToRaw return AVP struct of this value
func (v *MSCSMDeliveryOutcome) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3318, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	if v != nil {
		t := []dia.RawAVP{
			v.SMDeliveryCause.ToRaw(),
			v.AbsentUserDiagnosticSM.ToRaw()}
		a.Encode(dia.GroupedAVP(t))
	}
	return a
}

// SGSNSMDeliveryOutcome AVP shall indicate the outcome of
// the SM delivery for setting the message waiting data
// in the HSS when the SM delivery is with an SGSN.
type SGSNSMDeliveryOutcome struct {
	SMDeliveryCause        SMDeliveryCause
	AbsentUserDiagnosticSM AbsentUserDiagnosticSM
}

// ToRaw return AVP struct of this value
func (v *SGSNSMDeliveryOutcome) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3319, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	if v != nil {
		t := []dia.RawAVP{
			v.SMDeliveryCause.ToRaw(),
			v.AbsentUserDiagnosticSM.ToRaw()}
		a.Encode(dia.GroupedAVP(t))
	}
	return a
}

// IPSMGWSMDeliveryOutcome AVP shall indicate the outcome of
// the SM delivery for setting the message waiting data
// when the SM delivery is with an IP-SM-GW.
type IPSMGWSMDeliveryOutcome struct {
	SMDeliveryCause        SMDeliveryCause
	AbsentUserDiagnosticSM AbsentUserDiagnosticSM
}

// ToRaw return AVP struct of this value
func (v *IPSMGWSMDeliveryOutcome) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3320, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	if v != nil {
		t := []dia.RawAVP{
			v.SMDeliveryCause.ToRaw(),
			v.AbsentUserDiagnosticSM.ToRaw()}
		a.Encode(dia.GroupedAVP(t))
	}
	return a
}

// SMDeliveryCause AVP shall indicate the cause of
// the SMP delivery result.
type SMDeliveryCause dia.Enumerated

const (
	// UeMemoryCapacityExceeded is Enumerated value 0
	UeMemoryCapacityExceeded SMDeliveryCause = 0
	// AbsentUser is Enumerated value 1
	AbsentUser SMDeliveryCause = 1
	// SuccessfulTransfer is Enumerated value 2
	SuccessfulTransfer SMDeliveryCause = 2
)

// ToRaw return AVP struct of this value
func (v *SMDeliveryCause) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3321, VenID: 10415,
		FlgV: true, FlgM: true, FlgP: false}
	if v != nil {
		a.Encode(dia.Enumerated(*v))
	}
	return a
}

// FromRaw get AVP value
func (v *SMDeliveryCause) FromRaw(a dia.RawAVP) (e error) {
	if e = a.Validate(10415, 3321, true, true, false); e != nil {
		return
	}
	s := new(dia.Enumerated)
	if e = a.Decode(s); e != nil {
		return
	}
	*v = SMDeliveryCause(*s)
	return
}


// RDRFlags AVP contain a bit mask.
// SingleAttemptDelivery indicates that only one delivery attempt
// shall be performed for this particular SM.
type RDRFlags struct {
	SingleAttemptDelivery bool
}

// ToRaw return AVP struct of this value
func (v *RDRFlags) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3323, VenID: 10415,
		FlgV: true, FlgM: false, FlgP: false}
	if v != nil {
		i := uint32(0)
		if v.SingleAttemptDelivery {
			i = i | 0x00000001
		}
		a.Encode(i)
	}
	return a
}

// FromRaw get AVP value
func (v *RDRFlags) FromRaw(a dia.RawAVP) (e error) {
	if e = a.Validate(10415, 3323, true, false, false); e != nil {
		return
	}
	s := new(uint32)
	if e = a.Decode(s); e != nil {
		return
	}
	*v = RDRFlags{
		SingleAttemptDelivery: (*s)&0x00000001 == 0x00000001}
	return
}

// MaximumUEAvailabilityTime AVP shall contain the timestamp (in UTC)
// until which a UE using a power saving mechanism
// (such as extended idle mode DRX) is expected to be reachable
// for SM Delivery.
type MaximumUEAvailabilityTime time.Time

// ToRaw return AVP struct of this value
func (v *MaximumUEAvailabilityTime) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3329, VenID: 10415,
		FlgV: true, FlgM: false, FlgP: false}
	if v != nil {
		a.Encode(time.Time(*v))
	}
	return a
}

// FromRaw get AVP value
func (v *MaximumUEAvailabilityTime) FromRaw(a dia.RawAVP) (e error) {
	if e = a.Validate(10415, 3329, true, false, false); e != nil {
		return
	}
	s := new(time.Time)
	if e = a.Decode(s); e != nil {
		return
	}
	*v = MaximumUEAvailabilityTime(*s)
	return
}

// SMSGMSCAlertEvent AVP shall contain a bit mask.
// UEAvailableForMTSMS shall indicate that the UE is
// now available for MT SMS
// UEUnderNewServingNode shall indicate that the UE has moved
// under the coverage of another MME or SGSN.
type SMSGMSCAlertEvent struct {
	UEAvailableForMTSMS   bool
	UEUnderNewServingNode bool
}

// ToRaw return AVP struct of this value
func (v *SMSGMSCAlertEvent) ToRaw() dia.RawAVP {
	a := dia.RawAVP{Code: 3333, VenID: 10415,
		FlgV: true, FlgM: false, FlgP: false}
	if v != nil {
		i := uint32(0)
		if v.UEAvailableForMTSMS {
			i = i | 0x00000001
		}
		if v.UEUnderNewServingNode {
			i = i | 0x00000002
		}
		a.Encode(i)
	}
	return a
}

// FromRaw get AVP value
func (v *SMSGMSCAlertEvent) FromRaw(a dia.RawAVP) (e error) {
	if e = a.Validate(10415, 3333, true, false, false); e != nil {
		return
	}
	s := new(uint32)
	if e = a.Decode(s); e != nil {
		return
	}
	*v = SMSGMSCAlertEvent{
		UEAvailableForMTSMS:   (*s)&0x00000001 == 0x00000001,
		UEUnderNewServingNode: (*s)&0x00000002 == 0x00000002}
	return
}
*/
