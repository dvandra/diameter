package ts29338

/*
 <SRR> ::= < Diameter Header: 8388647, REQ, PXY, 16777312 >
           < Session-Id >
		   [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
           [ Destination-Host ]
           { Destination-Realm }
           [ MSISDN ]
           [ User-Name ]
           [ SMSMI-Correlation-ID ]
         * [ Supported-Features ]
           [ SC-Address ]
           [ SM-RP-MTI ]
           [ SM-RP-SMEA ]
           [ SRR-Flags ]
           [ SM-Delivery-Not-Intended ]
         * [ AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]

 <SRA> ::= < Diameter Header: 8388647, PXY, 16777312 >
           < Session-Id >
		   [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           [ Result-Code ]
           [ Experimental-Result ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
           [ User-Name ]
         * [ Supported-Features ]
           [ Serving-Node ]
           [ Additional-Serving-Node ]
           [ LMSI ]
           [ User-Identifier ]
           [ MWD-Status ]
           [ MME-Absent-User-Diagnostic-SM ]
           [ MSC-Absent-User-Diagnostic-SM ]
           [ SGSN-Absent-User-Diagnostic-SM ]
         * [ AVP ]
         * [ Failed-AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]
*/

/*
 <ALR> ::= < Diameter Header: 8388648, REQ, PXY, 16777312 >
           < Session-Id >
           [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
           [ Destination-Host ]
           { Destination-Realm }
           { SC-Address }
           { User-Identifier }
           [ SMSMI-Correlation-ID ]
           [ Maximum-UE-Availability-Time ]
           [ SMS-GMSC-Alert-Event ]
           [ Serving-Node ]
         * [ Supported-Features ]
         * [ AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]
 <ALA> ::= < Diameter Header: 8388648, PXY, 16777312 >
           < Session-Id >
           [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           [ Result-Code ]
           [ Experimental-Result ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
         * [ Supported-Features ]
         * [ AVP ]
         * [ Failed-AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]
*/

/*
 <RDR> ::= < Diameter Header: 8388649, REQ, PXY, 16777312 >
           < Session-Id >
           [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
           [ Destination-Host ]
           { Destination-Realm }
         * [ Supported-Features ]
           { User-Identifier }
           [ SMSMI-Correlation-ID ]
           { SC-Address }
           { SM-Delivery-Outcome }
           [ RDR-Flags ]
         * [ AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]
 <RDA> ::= < Diameter Header: 8388649, PXY, 16777312 >
           < Session-Id >
           [ DRMP ]
           [ Vendor-Specific-Application-Id ]
           [ Result-Code ]
           [ Experimental-Result ]
           { Auth-Session-State }
           { Origin-Host }
           { Origin-Realm }
         * [ Supported-Features ]
           [ User-Identifier ]
         * [ AVP ]
         * [ Failed-AVP ]
         * [ Proxy-Info ]
         * [ Route-Record ]
*/
