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

	// Vendor found

	if host.Vendor != "" {
		score += 15
	}

	// Hostname found

	if host.Hostname != "" {
		score += 15
	}

	// Open ports found

	if len(host.Ports) > 0 {
		score += 10
	}

	// HTTP information found

	if len(host.HTTP) > 0 {
		score += 10
	}

	// UDP service found

	if len(host.UDPServices) > 0 {
		score += 10
	}

	if score > 100 {
		score = 100
	}

	return score
}


func CollectEvidence(host models.Host) []string {

	evidence := []string{}

	for _, source := range host.Sources {

		switch source.Type {

		case models.DiscoveryARP:
			evidence = append(evidence, "ARP")

		case models.DiscoveryReverseDNS:
			evidence = append(evidence, "ReverseDNS")

		case models.DiscoveryNetBIOS:
			evidence = append(evidence, "NetBIOS")

		case models.DiscoveryMDNS:
			evidence = append(evidence, "mDNS")
		}
	}

	if host.MAC != "" {
		evidence = append(evidence, "MAC")
	}

	if host.Vendor != "" {
		evidence = append(evidence, "Vendor")
	}

	if host.Hostname != "" {
		evidence = append(evidence, "Hostname")
	}

	if len(host.Ports) > 0 {
		evidence = append(evidence, "TCP")
	}

	if len(host.HTTP) > 0 {
		evidence = append(evidence, "HTTP")
	}

	if len(host.UDPServices) > 0 {
		evidence = append(evidence, "UDP")
	}

	return evidence
}