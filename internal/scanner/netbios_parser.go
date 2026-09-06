package scanner

import (
	"encoding/binary"
	"strings"

	"OrsoNetwork/internal/logger"
)

func ParseNetBIOSResponse(data []byte) NetBIOSResult {

	result := NetBIOSResult{}

	// NetBIOS Name Service использует DNS-подобный
	// 12-байтовый заголовок.
	if len(data) < 12 {
		return result
	}

	transactionID := binary.BigEndian.Uint16(data[0:2])
	flags := binary.BigEndian.Uint16(data[2:4])
	questions := binary.BigEndian.Uint16(data[4:6])
	answers := binary.BigEndian.Uint16(data[6:8])
	authority := binary.BigEndian.Uint16(data[8:10])
	additional := binary.BigEndian.Uint16(data[10:12])

	logger.Debug(
		"NETBIOS HEADER:",
		"TRANSACTION:", transactionID,
		"FLAGS:", flags,
		"QUESTIONS:", questions,
		"ANSWERS:", answers,
		"AUTHORITY:", authority,
		"ADDITIONAL:", additional,
	)

	offset := 12

	if offset >= len(data) {
		return result
	}

	nameLength := int(data[offset])

	logger.Debug(
		"NETBIOS NAME LENGTH:",
		nameLength,
	)

	nameStart := offset + 1
	nameEnd := nameStart + nameLength

	if nameEnd > len(data) {
		return result
	}

	logger.Debug(
		"NETBIOS NAME:",
		data[nameStart:nameEnd],
	)

	offset = nameEnd

	if offset >= len(data) {
		return result
	}

	logger.Debug(
		"NETBIOS NAME TERMINATOR:",
		data[offset],
	)

	offset++

	// Дальше обязательно проверяем, что в пакете
	// достаточно байт для TYPE.
	if offset+2 > len(data) {
		return result
	}

	recordType := binary.BigEndian.Uint16(
		data[offset : offset+2],
	)

	logger.Debug(
		"NETBIOS RECORD TYPE:",
		recordType,
	)

	offset += 2

	// CLASS — ещё 2 байта.
	if offset+2 > len(data) {
		return result
	}

	recordClass := binary.BigEndian.Uint16(
		data[offset : offset+2],
	)

	logger.Debug(
		"NETBIOS RECORD CLASS:",
		recordClass,
	)

	offset += 2

	ttl := binary.BigEndian.Uint32(
		data[offset : offset+4],
	)

	logger.Debug(
		"NETBIOS RECORD TTL:",
		ttl,
	)

	offset += 4

	// RDLENGTH — размер RDATA в байтах.
	if offset+2 > len(data) {
		return result
	}

	rdLength := int(
		binary.BigEndian.Uint16(
			data[offset : offset+2],
		),
	)

	logger.Debug(
		"NETBIOS RDLENGTH:",
		rdLength,
	)

	offset += 2

	// RDATA
	if offset+rdLength > len(data) {
		return result
	}

	rdata := data[offset : offset+rdLength]

	if len(rdata) < 1 {
		return result
	}

	nameCount := int(rdata[0])

	logger.Debug(
		"NETBIOS NAME COUNT:",
		nameCount,
	)

	offset = 1

	for i := 0; i < nameCount; i++ {

		if offset+18 > len(rdata) {
			return result
		}

		nameBytes := rdata[offset : offset+15]

		suffix := rdata[offset+15]

		name := strings.TrimSpace(
			string(nameBytes),
		)

		flags := binary.BigEndian.Uint16(
			rdata[offset+16 : offset+18],
		)

		logger.Debug(
			"NETBIOS NAME ENTRY:",
			"INDEX:", i,
			"NAME:", name,
			"SUFFIX:", suffix,
			"FLAGS:", flags,
		)

		result.Names = append(
			result.Names,
			NetBIOSName{
				Name:   name,
				Suffix: suffix,
				Flags:  flags,
			},
		)

		offset += 18
	}

	logger.Debug(
		"NETBIOS RDATA OFFSET:",
		offset,
		"REMAINING:",
		len(rdata)-offset,
	)

	if len(rdata)-offset >= 6 {
		mac := rdata[offset : offset+6]

		logger.Debug(
			"NETBIOS UNIT ID:",
			mac,
		)
	}

	logger.Debug(
		"NETBIOS RDATA LENGTH:",
		len(rdata),
	)

	return result
}

func validNetBIOSMAC(mac string) bool {

	if mac == "" {
		return false
	}

	if mac == "00-00-00-00-00-00" {
		return false
	}

	return true
}
