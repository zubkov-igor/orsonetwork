package scanner

import "fmt"

func readLength(
	data []byte,
	offset *int,
) (int, error) {

	if *offset >= len(data) {
		return 0, fmt.Errorf("unexpected end of data")
	}

	first := data[*offset]
	(*offset)++

	// Short form.
	//
	// 0xxxxxxx
	//
	// The first byte directly contains
	// the length.
	if first&0x80 == 0 {
		return int(first), nil
	}

	// Long form.
	//
	// 1xxxxxxx
	//
	// The lower 7 bits tell us
	// how many bytes contain the length.
	count := int(first & 0x7F)

	if count == 0 {
		return 0, fmt.Errorf("indefinite length is not supported")
	}

	if *offset+count > len(data) {
		return 0, fmt.Errorf("invalid length field")
	}

	length := 0

	for i := 0; i < count; i++ {
		length <<= 8
		length |= int(data[*offset])
		(*offset)++
	}

	return length, nil
}

func readInteger(
	data []byte,
	offset *int,
) (int, error) {

	if *offset >= len(data) {
		return 0, fmt.Errorf("unexpected end of data")
	}

	tag := data[*offset]
	(*offset)++

	if tag != 0x02 {
		return 0, fmt.Errorf("expected INTEGER, got 0x%02x", tag)
	}

	length, err := readLength(data, offset)
	if err != nil {
		return 0, err
	}

	if *offset+length > len(data) {
		return 0, fmt.Errorf("invalid INTEGER length")
	}

	value := 0

	for i := 0; i < length; i++ {
		value <<= 8
		value |= int(data[*offset])
		(*offset)++
	}

	return value, nil
}

func readOctetString(
	data []byte,
	offset *int,
) (string, error) {

	if *offset >= len(data) {
		return "", fmt.Errorf("unexpected end of data")
	}

	tag := data[*offset]
	(*offset)++

	if tag != 0x04 {
		return "", fmt.Errorf(
			"expected OCTET STRING, got 0x%02x",
			tag,
		)
	}

	length, err := readLength(
		data,
		offset,
	)

	if err != nil {
		return "", err
	}

	if *offset+length > len(data) {
		return "", fmt.Errorf(
			"invalid OCTET STRING length",
		)
	}

	value := data[*offset : *offset+length]

	(*offset) += length

	return string(value), nil

}

func readOID(
	data []byte,
	offset *int,
) ([]int, error) {

	if *offset >= len(data) {
		return nil, fmt.Errorf("unexpected end of data")
	}

	tag := data[*offset]
	(*offset)++

	if tag != 0x06 {
		return nil, fmt.Errorf(
			"expected OBJECT IDENTIFIER, got 0x%02x",
			tag,
		)
	}

	length, err := readLength(
		data,
		offset,
	)

	if err != nil {
		return nil, err
	}

	if *offset+length > len(data) {
		return nil, fmt.Errorf(
			"invalid OBJECT IDENTIFIER length",
		)
	}

	end := *offset + length

	if *offset >= end {
		return nil, fmt.Errorf(
			"empty OBJECT IDENTIFIER",
		)
	}

	first := data[*offset]
	(*offset)++

	var firstArc int
	var secondArc int

	switch {
	case first < 40:
		firstArc = 0
		secondArc = int(first)

	case first < 80:
		firstArc = 1
		secondArc = int(first - 40)

	default:
		firstArc = 2
		secondArc = int(first - 80)
	}

	oid := []int{
		firstArc,
		secondArc,
	}

	value := 0

	for *offset < end {

		b := data[*offset]
		(*offset)++

		value <<= 7
		value |= int(b & 0x7F)

		if b&0x80 == 0 {
			oid = append(
				oid,
				value,
			)

			value = 0
		}
	}

	if value != 0 {
		return nil, fmt.Errorf(
			"unterminated OBJECT IDENTIFIER",
		)
	}

	return oid, nil
}
