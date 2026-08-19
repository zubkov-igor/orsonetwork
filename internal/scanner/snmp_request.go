package scanner

// BuildSNMPRequest builds SNMPv1 GET request.
//
// Request:
//   - Version: SNMPv1 (0)
//   - Community: public
//   - PDU: GetRequest
//   - VarBind: OID + NULL
//
// Used for querying values like:
// 1.3.6.1.2.1.1.1.0 (sysDescr)

func BuildSNMPRequest(
	oid []int,
) []byte {

	// VarBind:
	// SEQUENCE {
	//     OBJECT IDENTIFIER,
	//     NULL
	// }

	varBind := EncodeVarBind(
		oid,
	)

	// VarBind list:
	// SEQUENCE {
	//     VarBind
	// }

	varBindList := EncodeSequence(
		varBind,
	)

	// GET REQUEST PDU
	//
	// PDU {
	//     request-id
	//     error
	//     error-index
	//     varbind-list
	// }

	pduData := []byte{}

	pduData = append(
		pduData,
		EncodeInteger(1)...,
	)

	pduData = append(
		pduData,
		EncodeInteger(0)...,
	)

	pduData = append(
		pduData,
		EncodeInteger(0)...,
	)

	pduData = append(
		pduData,
		varBindList...,
	)

	pdu := EncodePDU(
		PDUGetRequest,
		pduData,
	)

	// SNMP message:
	//
	// SEQUENCE {
	//     version
	//     community
	//     PDU
	// }

	message := []byte{

		// version INTEGER 0
		0x02,
		0x01,
		0x00,

		// community OCTET STRING
		0x04,
		byte(len(SNMPCommunity)),
	}

	message = append(
		message,
		[]byte(SNMPCommunity)...,
	)

	message = append(
		message,
		pdu...,
	)

	return EncodeSequence(
		message,
	)
}
