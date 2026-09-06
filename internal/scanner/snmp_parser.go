package scanner

import (
	"fmt"

	"OrsoNetwork/internal/logger"
)

type SNMPResponse struct {
	OID   []int
	Value string
}

func ParseSNMPResponse(
	data []byte,
) SNMPResponse {

	logger.Debug(
		"SNMP PARSER CALLED:",
		len(data),
	)

	response := SNMPResponse{}

	offset := 0

	// SNMP message
	//
	// SEQUENCE {
	//     version
	//     community
	//     PDU
	// }

	if offset >= len(data) {
		return response
	}

	tag := data[offset]
	offset++

	if tag != 0x30 {
		return response
	}

	length, err := readLength(
		data,
		&offset,
	)

	if err != nil {
		return response
	}

	if offset+length > len(data) {
		return response
	}

	end := offset + length

	// Version
	//
	// INTEGER
	//
	// SNMPv1  = 0
	// SNMPv2c = 1

	version, err := readInteger(
		data,
		&offset,
	)

	if err != nil {
		return response
	}

	logger.Debug(
		"SNMP VERSION:",
		version,
	)

	// Community
	//
	// OCTET STRING

	community, err := readOctetString(
		data,
		&offset,
	)

	if err != nil {
		return response
	}

	logger.Debug(
		"SNMP COMMUNITY:",
		community,
	)

	if offset >= end {
		return response
	}

	// PDU

	tag = data[offset]
	offset++

	logger.Debug(
		"SNMP PDU TAG:",
		fmt.Sprintf("0x%02x", tag),
	)

	return response
}
