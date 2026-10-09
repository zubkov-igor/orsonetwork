package scanner

import (
	"time"

	"strconv"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichHosts(
	hosts []models.Host,
	iface models.Interface,
	config models.ScannerConfig,
) []models.Host {

	logger.Info(
		"HOSTS BEFORE ENRICHMENT:",
		len(hosts),
	)

	// =========================
	// mDNS
	// =========================

	if config.EnableMDNS {

		stepStart := time.Now()

		hosts = EnrichMDNS(
			hosts,
		)

		logger.Debug(
			"ENRICH TIMING MDNS:",
			time.Since(stepStart),
		)
	}

	// =========================
	// UDP
	// =========================

	if config.EnableUDP {

		stepStart := time.Now()

		hosts = EnrichUDP(
			hosts,
			iface,
			config,
		)

		logger.Debug(
			"ENRICH TIMING UDP:",
			time.Since(stepStart),
		)
	}

	// =========================
	// NetBIOS
	// =========================

	if config.EnableNetBIOS {

		stepStart := time.Now()

		hosts = EnrichNetBIOS(
			hosts,
		)

		logger.Debug(
			"ENRICH TIMING NETBIOS:",
			time.Since(stepStart),
		)
	}

	// =========================
	// SNMP
	// =========================

	if config.EnableSNMP {

		stepStart := time.Now()

		hosts = EnrichSNMP(
			hosts,
		)

		logger.Debug(
			"ENRICH TIMING SNMP:",
			time.Since(stepStart),
		)
	}

	// =========================
	// TCP Ports
	// =========================

	if config.EnableTCP {

		stepStart := time.Now()

		hosts = EnrichPorts(
			hosts,
		)

		logger.Debug(
			"ENRICH TIMING PORTS:",
			time.Since(stepStart),
		)
	}

	// =========================
	// SMB
	// =========================

	if config.EnableSMB {

		stepStart := time.Now()

		for i := range hosts {

			hasSMB := false

			for _, port := range hosts[i].Ports {

				if port.Number == 445 {

					hasSMB = true
					break
				}
			}

			if !hasSMB {
				continue
			}

			logger.Debug(
				"SMB ENRICH:",
				hosts[i].IP,
			)

			result := ProbeSMB(
				hosts[i].IP,
			)

			if !result.Found {
				continue
			}

			for _, share := range result.Shares {

				logger.Info(
					"SMB SHARE:",
					hosts[i].IP,
					"NAME:",
					share.Name,
					"TYPE:",
					share.Type,
					"COMMENT:",
					share.Comment,
				)

				hosts[i].SMBShares = append(
					hosts[i].SMBShares,
					models.SMBShare{
						Name:    share.Name,
						Type:    strconv.FormatUint(uint64(share.Type), 10),
						Comment: share.Comment,
					},
				)
			}

			logger.Info(
				"SMB SHARES:",
				hosts[i].IP,
				hosts[i].SMBShares,
			)
		}

		logger.Debug(
			"ENRICH TIMING SMB:",
			time.Since(stepStart),
		)
	}

	// =========================
	// Host Status
	// =========================

	for i := range hosts {

		UpdateHostStatus(
			&hosts[i],
			config,
		)
	}

	// =========================
	// Fingerprint
	// =========================

	stepStart := time.Now()

	for i := range hosts {

		hosts[i].Fingerprint = BuildFingerprint(
			hosts[i],
		)

		logger.Debug(
			"FINGERPRINT:",
			hosts[i].IP,
			hosts[i].Fingerprint,
		)
	}

	logger.Debug(
		"ENRICH TIMING FINGERPRINT:",
		time.Since(stepStart),
	)

	// =========================
	// Identification
	// =========================

	stepStart = time.Now()

	for i := range hosts {

		evidence := BuildFingerprint(hosts[i])

		hosts[i].Fingerprint = evidence

		identification := IdentifyDevice(
			hosts[i],
			evidence,
		)

		hosts[i].Type = identification.Type

		hosts[i].Confidence =
			CalculateConfidence(
				hosts[i],
			)

		logger.Debug(
			"IDENTIFICATION:",
			hosts[i].IP,
			"TYPE:",
			hosts[i].Type,
			"CONFIDENCE:",
			hosts[i].Confidence,
		)

		if hosts[i].Type != "unknown" {

			logger.Info(
				"DEVICE IDENTIFIED:",
				hosts[i].IP,
				"TYPE:",
				hosts[i].Type,
				"CONFIDENCE:",
				hosts[i].Confidence,
			)
		}
	}

	logger.Debug(
		"ENRICH TIMING IDENTIFICATION:",
		time.Since(stepStart),
	)

	return hosts
}

func UpdateHostStatus(
	host *models.Host,
	config models.ScannerConfig,
) {

	// ICMP

	if config.EnableICMP && host.Online {

		logger.Debug(
			"HOST STATUS:",
			host.IP,
			"ONLINE",
			"REASON: ICMP",
		)

		return
	}

	// ARP

	if config.EnableARP && host.MAC != "" {

		host.Online = true

		logger.Debug(
			"HOST STATUS:",
			host.IP,
			"ONLINE",
			"REASON: ARP",
		)

		return
	}

	// mDNS

	if config.EnableMDNS && len(host.MDNS) > 0 {

		host.Online = true

		logger.Debug(
			"HOST STATUS:",
			host.IP,
			"ONLINE",
			"REASON: mDNS SERVICE",
		)

		return
	}

	// UDP

	if config.EnableUDP && len(host.UDPServices) > 0 {

		host.Online = true

		logger.Debug(
			"HOST STATUS:",
			host.IP,
			"ONLINE",
			"REASON: UDP SERVICE",
		)

		return
	}

	// TCP

	if config.EnableTCP && len(host.Ports) > 0 {

		host.Online = true

		logger.Debug(
			"HOST STATUS:",
			host.IP,
			"ONLINE",
			"REASON: TCP PORT",
		)

		return
	}

	host.Online = false

	logger.Debug(
		"HOST STATUS:",
		host.IP,
		"OFFLINE",
		"REASON: NO RESPONSE",
	)

}
