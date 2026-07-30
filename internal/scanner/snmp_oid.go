package scanner

var (
	SNMPCommunity = "public"

	OIDSysDescr = []int{
		1,
		3,
		6,
		1,
		2,
		1,
		1,
		1,
		0,
	}
)

func EncodeOID(
	oid []int,
) []byte {

	if len(oid) < 2 {
		return nil
	}

	result := make(
		[]byte,
		0,
	)

	first := oid[0]*40 + oid[1]

	result = append(
		result,
		byte(first),
	)

	for _, value := range oid[2:] {

		result = append(
			result,
			encodeOIDValue(value)...,
		)
	}

	return result
}

func encodeOIDValue(
	value int,
) []byte {

	if value == 0 {
		return []byte{0}
	}

	var result []byte

	for value > 0 {

		result = append(
			[]byte{
				byte(value & 0x7F),
			},
			result...,
		)

		value >>= 7
	}

	for i := 0; i < len(result)-1; i++ {

		result[i] |= 0x80
	}

	return result
}
