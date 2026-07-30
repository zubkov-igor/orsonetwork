package scanner

func BuildSNMPRequest(
	oid []byte,
) []byte {

	// OID
	oidField := WrapBER(
		0x06,
		oid,
	)

	// NULL value
	nullField := []byte{
		0x05,
		0x00,
	}

	// VarBind
	varBind := append(
		oidField,
		nullField...,
	)

	varBind = WrapBER(
		0x30,
		varBind,
	)

	// VarBind list
	varBindList := WrapBER(
		0x30,
		varBind,
	)

	// GET REQUEST PDU
	pdu := []byte{

		0x02, 0x04,
		0x00, 0x00, 0x00, 0x01,

		0x02, 0x01, 0x00,

		0x02, 0x01, 0x00,
	}

	pdu = append(
		pdu,
		varBindList...,
	)

	pdu = WrapBER(
		0xA0,
		pdu,
	)

	// SNMP message body

	message := []byte{

		0x02, 0x01, 0x00,

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

	return WrapBER(
		0x30,
		message,
	)

}
