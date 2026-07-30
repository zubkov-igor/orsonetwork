package scanner

import "OrsoNetwork/internal/models"

func CalculateConfidence(
	host models.Host,
) int {

	score := 0

	for _, source := range host.Sources {

		switch source.Type {

		case models.DiscoveryARP:
			score += 25

		case models.DiscoveryReverseDNS:
			score += 15

		case models.DiscoveryNetBIOS:
			score += 25

		case models.DiscoveryMDNS:
			score += 20

		case models.DiscoveryTCP:
			score += 10

		case models.DiscoveryUDP:
			score += 10
			
		case models.DiscoverySNMP:
    		score += 30

		case models.DiscoverySSDP:
    		score += 20
		}
	}


	// MAC сильно повышает уверенность,
	// потому что это физический идентификатор устройства
	if host.MAC != "" {
		score += 10
	}


	// Открытые порты помогают определить устройство
	if len(host.Ports) > 0 {
		score += 10
	}


	if len(host.UDPServices) > 0 {
		score += 5
	}


	// Ограничение сверху
	if score > 100 {
		score = 100
	}


	return score
}