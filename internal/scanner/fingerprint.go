package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func BuildFingerprint(host models.Host) []models.FingerprintEvidence {

	var evidence []models.FingerprintEvidence

	// =========================
	// TCP
	// =========================

	logger.Debug(
		"TCP FINGERPRINT INPUT:",
		host.IP,
		host.Ports,
	)

	for _, port := range host.Ports {

		switch port.Number {

		case 22:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "ssh",
			})
		case 135:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "msrpc",
			})

		case 139:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "netbios",
			})

		case 445:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "smb",
			})

		case 3389:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "rdp",
			})
		case 7680:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "port:7680",
			})

		case 554:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "rtsp",
			})

		case 631:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "ipp",
			})
		case 9100:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "raw-printing",
			})
		}
	}

	// =========================
	// mDNS
	// =========================

	logger.Debug(
		"MDNS FINGERPRINT INPUT:",
		host.IP,
		host.MDNS,
	)

	for _, service := range host.MDNS {

		evidence = append(
			evidence,
			models.FingerprintEvidence{
				Source: models.DiscoveryMDNS,
				Value:  service.Service,
			},
		)
	}
	for _, httpInfo := range host.HTTP {

		if httpInfo.Server != "" {
			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryTCP,
					Value:  "http-server:" + httpInfo.Server,
				},
			)
		}

		if httpInfo.Title != "" {
			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryTCP,
					Value:  "http-title:" + httpInfo.Title,
				},
			)
		}

		for _, fingerprint := range httpInfo.Fingerprint {
			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryTCP,
					Value:  "http-fingerprint:" + fingerprint,
				},
			)
		}
	}

	for _, source := range host.Sources {
		if source.Type == models.DiscoveryNetBIOS &&
			source.Value != "" {

			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryNetBIOS,
					Value:  "name:" + source.Value,
				},
			)
		}
	}

	return evidence
}
