package scanner

func EncodeLength(length int) []byte {

	if length < 128 {
		return []byte{
			byte(length),
		}
	}

	bytes := []byte{}

	for length > 0 {
		bytes = append(
			[]byte{byte(length & 0xff)},
			bytes...,
		)

		length >>= 8
	}

	return append(
		[]byte{0x80 | byte(len(bytes))},
		bytes...,
	)
}

func WrapBER(
	tag byte,
	data []byte,
) []byte {

	result := []byte{
		tag,
	}

	result = append(
		result,
		EncodeLength(len(data))...,
	)

	result = append(
		result,
		data...,
	)

	return result
}

func EncodeInteger(value int) []byte {
	return WrapBER(
		0x02,
		[]byte{
			byte(value),
		},
	)

	var result []byte

	for value > 0 {
		result = append(
			[]byte{byte(value & 0xFF)},
			result...,
		)

		value >>= 8
	}

	// INTEGER is signed in BER.
	// If the highest bit is 1,
	// prepend 0x00 to keep the number positive.
	if result[0]&0x80 != 0 {
		result = append(
			[]byte{0x00},
			result...,
		)
	}

	return WrapBER(
		0x02,
		result,
	)
}

func EncodeOctetString(
	data []byte,
) []byte {

	return WrapBER(
		0x04,
		data,
	)

}

func EncodeNull() []byte {

	return []byte{
		0x05,
		0x00,
	}

}

func EncodeSequence(
	data []byte,
) []byte {

	return WrapBER(
		0x30,
		data,
	)

}
