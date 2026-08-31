package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
	"sync"
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
	config models.ScannerConfig,
) []models.UDPService {

	logger.Debug(
		"UDP DISCOVERY START:",
		ip,
	)

	var services []models.UDPService

	// =========================
	// NetBIOS UDP/137
	// =========================

	if config.EnableNetBIOS {

		data, err := ProbeNetBIOS(ip)

		if err == nil {

			logger.Info(
				"UDP NETBIOS RESPONSE:",
				ip,
				len(data),
			)

			netbios := ParseNetBIOSResponse(data)

			if netbios.Name != "" {

				logger.Info(
					"UDP NETBIOS FOUND:",
					ip,
					netbios.Name,
				)

				services = append(
					services,
					models.UDPService{
						IP:       ip,
						Port:     137,
						Service:  "NetBIOS",
						Protocol: "udp",
						Info:     netbios.Name,
					},
				)
			}
		}
	}

	logger.Debug(
		"UDP DISCOVERY FINISHED:",
		ip,
		len(services),
	)

	return services
}

func discoveryWorker(
	jobs <-chan string,
	results chan<- models.Host,
	wg *sync.WaitGroup,
	config models.ScannerConfig,
) {

	defer wg.Done()

	for ip := range jobs {

		host := discoverHostNew(
			ip,
			config,
		)

		results <- host
	}
}
