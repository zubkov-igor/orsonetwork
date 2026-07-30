package scanner

func EncodeLength(
	length int,
) []byte {

	if length < 128 {

		return []byte{
			byte(length),
		}

	}

	return []byte{
		0x81,
		byte(length),
	}

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

func EncodeInteger(
	value byte,
) []byte {

	return WrapBER(
		0x02,
		[]byte{
			value,
		},
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
