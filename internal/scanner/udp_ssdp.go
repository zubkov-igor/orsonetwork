package scanner

import (
	"net"
	"strings"
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

type SSDPResponse struct {
	IP       string
	Location string
	Server   string
	ST       string
	USN      string
}

func ProbeSSDP(
	iface models.Interface,
) []SSDPResponse {

	localAddr := &net.UDPAddr{
		IP:   net.ParseIP(iface.IP),
		Port: 0,
	}

	conn, err := net.ListenUDP(
		"udp4",
		localAddr,
	)

	if err != nil {

		logger.Log.Println(
			"SSDP CONNECT ERROR:",
			err,
		)

		return nil
	}

	defer conn.Close()

	logger.Log.Println(
		"SSDP LOCAL:",
		conn.LocalAddr(),
	)

	target := &net.UDPAddr{
		IP:   net.ParseIP("239.255.255.250"),
		Port: 1900,
	}

	request :=
		"M-SEARCH * HTTP/1.1\r\n" +
			"HOST: 239.255.255.250:1900\r\n" +
			"MAN: \"ssdp:discover\"\r\n" +
			"MX: 2\r\n" +
			"ST: ssdp:all\r\n" +
			"\r\n"

	n, err := conn.WriteToUDP(
		[]byte(request),
		target,
	)

	if err != nil {

		logger.Log.Println(
			"SSDP WRITE ERROR:",
			err,
		)

		return nil
	}

	logger.Log.Println(
		"SSDP SENT:",
		n,
		"bytes",
		target,
	)

	buffer := make([]byte, 4096)

	deadline := time.Now().Add(
		3 * time.Second,
	)

	var responses []SSDPResponse

	for {

		err = conn.SetReadDeadline(
			deadline,
		)

		if err != nil {
			break
		}

		n, addr, err := conn.ReadFromUDP(
			buffer,
		)

		if err != nil {

			logger.Log.Println(
				"SSDP READ FINISHED:",
				err,
			)

			break
		}

		response := string(
			buffer[:n],
		)

		logger.Log.Println(
			"SSDP RESPONSE FROM:",
			addr.IP,
		)

		logger.Log.Println(
			"SSDP RESPONSE:",
			response,
		)

		headers := parseSSDPHeaders(
			response,
		)

		responses = append(
			responses,
			SSDPResponse{
				IP:       addr.IP.String(),
				Location: headers["location"],
				Server:   headers["server"],
				ST:       headers["st"],
				USN:      headers["usn"],
			},
		)
	}

	logger.Log.Println(
		"SSDP DISCOVERY FOUND:",
		len(responses),
	)

	return responses
}

func parseSSDPHeaders(
	response string,
) map[string]string {

	headers := make(
		map[string]string,
	)

	lines := strings.Split(
		response,
		"\r\n",
	)

	for _, line := range lines {

		parts := strings.SplitN(
			line,
			":",
			2,
		)

		if len(parts) != 2 {
			continue
		}

		key := strings.ToLower(
			strings.TrimSpace(parts[0]),
		)

		value := strings.TrimSpace(
			parts[1],
		)

		headers[key] = value
	}

	return headers
}