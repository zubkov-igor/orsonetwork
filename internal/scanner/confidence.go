package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func CalculateConfidence(
	host models.Host,
	evidence []models.FingerprintEvidence,
	deviceType models.DeviceType,
	scores map[models.DeviceType]int,
) int {

	if deviceType == models.DeviceUnknown {
		return 0
	}

	bestScore := scores[deviceType]
	secondScore := 0

	for candidateType, candidateScore := range scores {

		if candidateType == models.DeviceUnknown {
			continue
		}

		if candidateType == deviceType {
			continue
		}

		if candidateScore > secondScore {
			secondScore = candidateScore
		}
	}

	if bestScore < 20 {
		return 0
	}

	margin := bestScore - secondScore

	strength := float64(bestScore-20) / 80 * 100

	if strength > 100 {
		strength = 100
	}

	marginScore := float64(margin) / float64(bestScore) * 100

	confidence := int(
		strength * marginScore / 100,
	)

	if confidence > 100 {
		confidence = 100
	}

	logger.Debug(
		"CONFIDENCE INPUT:",
		host.IP,
		"TYPE:",
		deviceType,
		"BEST SCORE:",
		bestScore,
		"SECOND SCORE:",
		secondScore,
		"MARGIN:",
		margin,
		"STRENGTH:",
		strength,
		"MARGIN SCORE:",
		marginScore,
		"CONFIDENCE:",
		confidence,
	)

	return confidence
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
