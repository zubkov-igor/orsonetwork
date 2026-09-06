package scanner

import (
	"net"
	"strings"
	"time"

	"OrsoNetwork/internal/logger"
)

func ProbeNetBIOS(ip string) ([]byte, error) {

	conn, err := net.DialUDP(
		"udp",
		nil,
		&net.UDPAddr{
			IP:   net.ParseIP(ip),
			Port: 137,
		},
	)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	encodedName := []byte("CK" + strings.Repeat("CA", 15))

	query := []byte{
		0x7b, 0x9d,
		0x00, 0x00,
		0x00, 0x01,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,

		0x20,
	}

	query = append(query, encodedName...)

	query = append(query,
		0x00,

		0x00, 0x21,
		0x00, 0x01,
	)

	logger.Debug(
		"NETBIOS QUERY:",
		ip,
		query,
		"LEN:",
		len(query),
	)

	_, err = conn.Write(query)

	if err != nil {
		return nil, err
	}

	conn.SetReadDeadline(
		time.Now().Add(
			2 * time.Second,
		),
	)

	buf := make([]byte, 512)

	n, addr, err := conn.ReadFromUDP(buf)

	if err != nil {
		logger.Debug(
			"NETBIOS ERROR:",
			ip,
			err,
		)

		return nil, err
	}

	logger.Debug(
		"NETBIOS RESPONSE:",
		ip,
		"FROM:",
		addr,
		"BYTES:",
		n,
	)

	logger.Debug(
		"NETBIOS RESPONSE DATA:",
		ip,
		buf[:n],
	)

	return buf[:n], nil
}
