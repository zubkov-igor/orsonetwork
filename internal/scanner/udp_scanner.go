package scanner

import (
	"sync"
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

type UDPDiscoveryJob struct {
	Index int
	IP    string
}

type UDPDiscoveryResult struct {
	Index    int
	Services []models.UDPService
}

func DiscoverUDP(
	ip string,
	iface models.Interface,
) []models.UDPService {

	logger.Log.Println(
		"UDP DISCOVERY START:",
		ip,
	)

	var services []models.UDPService

	udpPorts := []int{
		137,
		161,
	}

	for _, port := range udpPorts {

		service := UDPServices[port]

		var result UDPProbeResult

		switch port {

		case 161:

			logger.Log.Println(
				"UDP SNMP PROBE:",
				ip,
			)

			result = ProbeSNMP(
				ip,
			)

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

			services = append(
				services,
				models.UDPService{
					IP:       ip,
					Port:     port,
					Service:  service,
					Protocol: "udp",
					Info:     result.Info,
				},
			)
		}
	}

	logger.Log.Println(
		"UDP DISCOVERY FINISHED:",
		ip,
		len(services),
	)

	return services
}

func udpDiscoveryWorker(
	jobs <-chan UDPDiscoveryJob,
	results chan<- UDPDiscoveryResult,
	iface models.Interface,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for job := range jobs {

		services := DiscoverUDP(
			job.IP,
			iface,
		)

		results <- UDPDiscoveryResult{
			Index:    job.Index,
			Services: services,
		}
	}
}