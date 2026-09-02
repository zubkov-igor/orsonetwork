package scanner

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"OrsoNetwork/internal/logger"
)

func GrabBanner(
	ip string,
	port int,
) string {

	address := net.JoinHostPort(
		ip,
		strconv.Itoa(port),
	)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		500*time.Millisecond,
	)

	if err != nil {
		return ""
	}

	defer conn.Close()

	service := DetectPortService(port)

	switch service {
	case "ssh":
		return GrabSSHBanner(conn)
	case "ftp":
		return GrabFTPBanner(conn)
	case "telnet":
		return GrabTelnetBanner(conn)
	case "rtsp":
		return GrabRTSPBanner(conn, ip)
	case "http":
		return GrabHTTPBanner(conn, ip)
	case "https":
		return GrabHTTPSBanner(conn, ip)
	case "ipp":
		return GrabIPPBanner(conn, ip)
	case "rdp":
		return GrabRDPBanner(conn)
	}

	err = conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)

	if err != nil {
		return ""
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		return ""
	}

	return string(buffer[:n])
}

func GrabHTTPBanner(
	conn net.Conn,
	ip string,
) string {

	request := fmt.Sprintf(
		"GET / HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"Connection: close\r\n"+
			"\r\n",
		ip,
	)

	_, err := conn.Write(
		[]byte(request),
	)

	if err != nil {
		return ""
	}

	err = conn.SetReadDeadline(
		time.Now().Add(2 * time.Second),
	)

	if err != nil {
		return ""
	}

	response, err := http.ReadResponse(
		bufio.NewReader(conn),
		nil,
	)

	if err != nil {
		return ""
	}

	defer response.Body.Close()

	var banner []string

	if response.Status != "" {
		banner = append(
			banner,
			response.Status,
		)
	}

	if server := response.Header.Get("Server"); server != "" {
		banner = append(
			banner,
			"Server: "+server,
		)
	}

	if poweredBy := response.Header.Get("X-Powered-By"); poweredBy != "" {
		banner = append(
			banner,
			"X-Powered-By: "+poweredBy,
		)
	}

	if location := response.Header.Get("Location"); location != "" {
		banner = append(
			banner,
			"Location: "+location,
		)
	}

	return strings.Join(
		banner,
		"\n",
	)
}

func GrabHTTPSBanner(
	conn net.Conn,
	ip string,
) string {

	logger.Debug(
		"HTTPS BANNER START:",
		ip,
	)

	tlsConn := tls.Client(
		conn,
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         ip,
		},
	)

	err := tlsConn.Handshake()
	if err != nil {
		logger.Debug(
			"HTTPS TLS HANDSHAKE ERROR:",
			ip,
			err,
		)
		return ""
	}

	logger.Debug(
		"HTTPS TLS HANDSHAKE OK:",
		ip,
	)

	defer tlsConn.Close()

	request := fmt.Sprintf(
		"GET / HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"Connection: close\r\n"+
			"\r\n",
		ip,
	)

	_, err = tlsConn.Write(
		[]byte(request),
	)

	if err != nil {
		logger.Debug(
			"HTTPS REQUEST WRITE ERROR:",
			ip,
			err,
		)
		return ""
	}

	err = tlsConn.SetReadDeadline(
		time.Now().Add(2 * time.Second),
	)

	if err != nil {
		logger.Debug(
			"HTTPS READ DEADLINE ERROR:",
			ip,
			err,
		)
		return ""
	}

	response, err := http.ReadResponse(
		bufio.NewReader(tlsConn),
		nil,
	)

	if err != nil {
		logger.Debug(
			"HTTPS HTTP RESPONSE ERROR:",
			ip,
			err,
		)
		return ""
	}

	defer response.Body.Close()

	var banner []string

	if response.Status != "" {
		banner = append(
			banner,
			response.Status,
		)
	}

	if server := response.Header.Get("Server"); server != "" {
		banner = append(
			banner,
			"Server: "+server,
		)
	}

	if poweredBy := response.Header.Get("X-Powered-By"); poweredBy != "" {
		banner = append(
			banner,
			"X-Powered-By: "+poweredBy,
		)
	}

	if location := response.Header.Get("Location"); location != "" {
		banner = append(
			banner,
			"Location: "+location,
		)
	}

	return strings.Join(
		banner,
		"\n",
	)
}

func GrabSSHBanner(
	conn net.Conn,
) string {

	err := conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)

	if err != nil {
		return ""
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabFTPBanner(
	conn net.Conn,
) string {

	err := conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)

	if err != nil {
		return ""
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabRTSPBanner(
	conn net.Conn,
	ip string,
) string {

	request := fmt.Sprintf(
		"OPTIONS rtsp://%s/ RTSP/1.0\r\n"+
			"CSeq: 1\r\n"+
			"\r\n",
		ip,
	)

	_, err := conn.Write(
		[]byte(request),
	)

	if err != nil {
		return ""
	}

	err = conn.SetReadDeadline(
		time.Now().Add(2 * time.Second),
	)

	if err != nil {
		return ""
	}

	buffer := make([]byte, 4096)

	n, err := conn.Read(buffer)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabIPPBanner(
	conn net.Conn,
	ip string,
) string {

	request := fmt.Sprintf(
		"GET / HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"Connection: close\r\n"+
			"\r\n",
		ip,
	)

	_, err := conn.Write([]byte(request))
	if err != nil {
		return ""
	}

	err = conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)
	if err != nil {
		return ""
	}

	buffer := make([]byte, 4096)

	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabPrinterBanner(
	conn net.Conn,
) string {

	err := conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)
	if err != nil {
		return ""
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabTelnetBanner(
	conn net.Conn,
) string {

	err := conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)

	if err != nil {
		return ""
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}

func GrabRDPBanner(
	conn net.Conn,
) string {

	request := []byte{
		0x03, 0x00, 0x00, 0x13,
		0x0e, 0xe0, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x08, 0x00,
		0x03, 0x00, 0x00, 0x00,
	}

	_, err := conn.Write(request)
	if err != nil {
		return ""
	}

	err = conn.SetReadDeadline(
		time.Now().Add(500 * time.Millisecond),
	)
	if err != nil {
		return ""
	}

	buffer := make([]byte, 4096)

	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(buffer[:n]),
	)
}
