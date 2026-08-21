package scanner

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"OrsoNetwork/internal/logger"
)

func ProbeSNMP(
	ip string,
) UDPProbeResult {

	addr := ip + ":161"

	logger.Log.Println(
		"SNMP CONNECT:",
		addr,
	)

	conn, err := net.DialTimeout(
		"udp",
		addr,
		2*time.Second,
	)

	if err != nil {

		logger.Log.Println(
			"SNMP CONNECT ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	defer conn.Close()

	logger.Log.Println(
		"OID:",
		OIDSysDescr,
	)

	request := BuildSNMPRequest(
		OIDSysDescr,
	)

	logger.Log.Println(
		"SNMP REQUEST HEX:",
		fmt.Sprintf("% X", request),
	)

	logger.Log.Println(
		"SNMP SEND:",
		addr,
	)

	_, err = conn.Write(request)

	if err != nil {

		logger.Log.Println(
			"SNMP WRITE ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	logger.Log.Println(
		"SNMP WAIT:",
		ip,
	)

	buffer := make(
		[]byte,
		2048,
	)

	err = conn.SetReadDeadline(
		time.Now().Add(
			2 * time.Second,
		),
	)

	if err != nil {

		logger.Log.Println(
			"SNMP DEADLINE ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	n, err := conn.Read(
		buffer,
	)

	if err != nil {

		logger.Log.Println(
			"SNMP READ ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	response := buffer[:n]

	logger.Log.Println(
		"SNMP RESPONSE SIZE:",
		len(response),
	)

	logger.Log.Println(
		"SNMP RESPONSE HEX:",
		fmt.Sprintf("% X", response),
	)

	logger.Log.Println(
		"SNMP RESPONSE HEX COMPACT:",
		hex.EncodeToString(response),
	)

	if len(response) == 0 {

		logger.Log.Println(
			"SNMP EMPTY RESPONSE:",
			ip,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	if response[0] != 0x30 {

		logger.Log.Println(
			"SNMP INVALID BER RESPONSE:",
			ip,
			response[0],
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	parsed := ParseSNMPResponse(
		response,
	)
	if err != nil {

		logger.Log.Println(
			"SNMP PARSE ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	logger.Log.Println(
		"SNMP OID:",
		parsed.OID,
	)

	logger.Log.Println(
		"SNMP VALUE:",
		parsed.Value,
	)

	return UDPProbeResult{
		Found: true,
		Info:  "SNMP response",
		Raw:   response,
	}
}
