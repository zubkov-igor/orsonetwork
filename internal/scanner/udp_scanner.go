package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func DiscoverUDP(ip string) []models.UDPService {
	logger.Log.Println("UDP DISCOVERY START:", ip)

	var services []models.UDPService

	udpPorts := []int{
		137,
		161,
		1900,
		5353,
	}

	for _, port := range udpPorts {

		service := UDPServices[port]

		var result UDPProbeResult

		switch port {
		case 1900:
			logger.Log.Println("UDP SSDP PROBE:", ip)
			result = ProbeSSDP(ip)

		case 5353:
			logger.Log.Println("UDP MDNS PROBE:", ip)
			result = ProbeMDNS(ip)

		case 161:
			logger.Log.Println("UDP SNMP PROBE:", ip)
			result = ProbeSNMP(ip)

		default:
			continue
		}

		if result.Found {
			logger.Log.Println(
				"UDP SERVICE FOUND:",
				ip,
				port,
				service,
			)

			services = append(services, models.UDPService{
				IP:       ip,
				Port:     port,
				Service:  service,
				Protocol: "udp",
				Info:     result.Info,
			})
		}
	}

	logger.Log.Println("UDP DISCOVERY FINISHED:", ip, len(services))
	return services
}
