package scanner

import (
	"net"
	"strings"
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func ProbeSSDP(
	ip string,
	iface models.Interface,
) UDPProbeResult {

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

		return UDPProbeResult{
			Found: false,
		}
	}

	defer conn.Close()

	logger.Log.Println(
		"SSDP LOCAL:",
		conn.LocalAddr(),
	)

	logger.Log.Println(
		"SSDP INTERFACE:",
		iface.Name,
		iface.IP,
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

		return UDPProbeResult{
			Found: false,
		}
	}

	logger.Log.Println(
		"SSDP SENT:",
		n,
		"bytes",
		target,
	)

	buffer := make([]byte, 2048)

	deadline := time.Now().Add(
		2 * time.Second,
	)

	for {

		err = conn.SetReadDeadline(
			deadline,
		)

		if err != nil {
			return UDPProbeResult{
				Found: false,
			}
		}

		n, addr, err := conn.ReadFromUDP(
			buffer,
		)

		if err != nil {

			logger.Log.Println(
				"SSDP READ ERROR:",
				err,
			)

			return UDPProbeResult{
				Found: false,
			}
		}

		logger.Log.Println(
			"SSDP RESPONSE FROM:",
			addr.IP,
		)

		response := strings.ToLower(
			string(buffer[:n]),
		)

		logger.Log.Println(
			"SSDP RESPONSE:",
			response,
		)

		if addr.IP.String() != ip {

			logger.Log.Println(
				"SSDP RESPONSE SKIP:",
				addr.IP,
				"EXPECTED:",
				ip,
			)

			continue
		}

		if strings.Contains(
			response,
			"200 ok",
		) {

			logger.Log.Println(
				"SSDP RESPONSE FOUND:",
				ip,
			)

			return UDPProbeResult{
				Found: true,
				Info:  response,
			}
		}
	}
}
