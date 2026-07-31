package scanner

import "OrsoNetwork/internal/models"

func CalculateConfidence(
	host models.Host,
) int {

	score := 0


	for _, source := range host.Sources {

		switch source.Type {

		case models.DiscoveryARP:
			score += 20

		case models.DiscoveryReverseDNS:
			score += 20

		case models.DiscoveryNetBIOS:
			score += 30

		case models.DiscoveryMDNS:
			score += 25

		}
	}


	// MAC address found

	if host.MAC != "" {
		score += 10
	}


	// Open ports found

	if len(host.Ports) > 0 {
		score += 10
	}


	// Device fingerprint confidence

	switch host.Type {

	case models.DeviceRouter:
		score += 15

	case models.DeviceGateway:
		score += 15

	case models.DeviceComputer:
		score += 10

	case models.DeviceServer:
		score += 15

	case models.DeviceCamera:
		score += 15

	case models.DevicePrinter:
		score += 15

	case models.DeviceNAS:
		score += 15
	}


	if score > 100 {
		score = 100
	}


	return score
}