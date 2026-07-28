package scanner

import (
    "fmt"
    "net"
    "time"
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func ScanPorts(
	ip string,
) []models.Port {

	var ports []models.Port

	logger.Log.Println(
    "PORT SCAN START:",
    ip,
)

	for _, port := range CommonPorts {

		if IsPortOpen(ip, port) {

			ports = append(
				ports,
				models.Port{
					Number:   port,
					Protocol: "tcp",
					Service:  DetectPortService(port),
					Open:     true,
				},
			)
		}
	}

	logger.Log.Println(
    "PORT SCAN FINISHED:",
    ip,
    len(ports),
)

	return ports
}

func IsPortOpen(
	ip string,
	port int,
) bool {

	address := fmt.Sprintf(
		"%s:%d",
		ip,
		port,
	)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		500*time.Millisecond,
	)

	if err != nil {
		return false
	}

	conn.Close()

	return true
}
