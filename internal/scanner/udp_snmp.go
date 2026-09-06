package scanner

import (
	"fmt"
	"net"
	"time"

	"OrsoNetwork/internal/logger"
)

func ProbeSNMP(
	ip string,
) UDPProbeResult {

	probeStart := time.Now()

	defer func() {

		logger.Debug(
			"SNMP PROBE DURATION:",
			ip,
			time.Since(probeStart),
		)
	}()

	addr := ip + ":161"

	logger.Debug(
		"SNMP CONNECT:",
		addr,
	)

	conn, err := net.DialTimeout(
		"udp",
		addr,
		500*time.Millisecond,
	)

	if err != nil {

		logger.Debug(
			"SNMP CONNECT ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	defer conn.Close()

	logger.Debug(
		"OID:",
		OIDSysDescr,
	)

	request := BuildSNMPRequest(
		OIDSysDescr,
	)

	logger.Debug(
		"SNMP REQUEST HEX:",
		fmt.Sprintf("% X", request),
	)

	logger.Debug(
		"SNMP SEND:",
		addr,
	)

	_, err = conn.Write(request)

	if err != nil {

		logger.Debug(
			"SNMP WRITE ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	logger.Debug(
		"SNMP WAIT:",
		ip,
	)

	buffer := make(
		[]byte,
		2048,
	)

	err = conn.SetReadDeadline(
		time.Now().Add(
			500 * time.Millisecond,
		),
	)

	if err != nil {

		logger.Debug(
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

	logger.Debug(
		"SNMP AFTER READ:",
		ip,
		"N:",
		n,
		"ERR:",
		err,
	)

	if err != nil {
		logger.Debug(
			"SNMP READ ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	if err != nil {

		logger.Debug(
			"SNMP READ ERROR:",
			ip,
			err,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	response := buffer[:n]

	logger.Debug(
		"SNMP RESPONSE SIZE:",
		len(response),
	)

	logger.Debug(
		"SNMP RESPONSE HEX:",
		fmt.Sprintf("% X", response),
	)

	if len(response) == 0 {

		logger.Debug(
			"SNMP EMPTY RESPONSE:",
			ip,
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	if response[0] != 0x30 {

		logger.Debug(
			"SNMP INVALID BER RESPONSE:",
			ip,
			response[0],
		)

		return UDPProbeResult{
			Found: false,
		}
	}

	logger.Debug(
		"SNMP BEFORE PARSER:",
		ip,
	)

	parsed := ParseSNMPResponse(
		response,
	)

	logger.Info(
		"SNMP FOUND:",
		ip,
		"OID:",
		parsed.OID,
		"VALUE:",
		parsed.Value,
	)

	return UDPProbeResult{
		Found: true,
		Info:  "SNMP response",
		Raw:   response,
	}

}
