package scanner

const (
	PDUGetRequest = 0xA0
	PDUGetNext    = 0xA1
	PDUResponse   = 0xA2
	PDUSetRequest = 0xA3
)

// EncodePDU wraps any SNMP PDU.
func EncodePDU(
	tag byte,
	data []byte,
) []byte {

	return WrapBER(
		tag,
		data,
	)
}

// EncodeVarBind builds
//
//	SEQUENCE {
//	    OBJECT IDENTIFIER
//	    NULL
//	}
//
// GET request always sends NULL value.
func EncodeVarBind(
	oid []int,
) []byte {

	oidBytes := EncodeOID(
		oid,
	)

	nullBytes := EncodeNull()

	data := append(
		oidBytes,
		nullBytes...,
	)

	return EncodeSequence(
		data,
	)
}

func EncodeVarBindList(
	oid []int,
) []byte {

	varBind := EncodeVarBind(
		oid,
	)

	return EncodeSequence(
		varBind,
	)
}

func EncodeGetRequest(
	requestID int,
	oid []int,
) []byte {

	requestIDBytes := EncodeInteger(
		requestID,
	)

	errorStatus := EncodeInteger(
		0,
	)

	errorIndex := EncodeInteger(
		0,
	)

	varBindList := EncodeVarBindList(
		oid,
	)

	data := append(
		requestIDBytes,
		errorStatus...,
	)

	data = append(
		data,
		errorIndex...,
	)

	data = append(
		data,
		varBindList...,
	)

	return EncodePDU(
		PDUGetRequest,
		data,
	)
}
