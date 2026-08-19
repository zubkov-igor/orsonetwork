package scanner

type SNMPResponse struct {
	OID   []int
	Value string
}

func ParseSNMPResponse(
	data []byte,
) SNMPResponse {

	response := SNMPResponse{}

	// BER sequence
	// version
	// community
	// PDU
	// varbind

	return response
}
