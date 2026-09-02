package scanner

import (
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func discoverHostNew(
	ip string,
	config models.ScannerConfig,
) models.Host {

	host := models.Host{
		IP: ip,
	}
	start := time.Now()

	// =========================
	// ICMP
	// =========================

	if config.EnableICMP {

		stepStart := time.Now()

		result := PingHost(
			ip,
			500*time.Millisecond,
		)

		pingDuration := time.Since(stepStart)

		if result.Online {

			host.Online = true
			host.RTT = result.RTT

			host.Sources = append(
				host.Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryICMP,
					Value: result.RTT.String(),
				},
			)
		}

		if pingDuration > 100*time.Millisecond {
			logger.Debug(
				"DISCOVERY SLOW:",
				ip,
				"PING:",
				pingDuration,
			)
		}
	}

	// ARP

	if config.EnableARP {

		mac := ARPResolve(ip)

		logger.Debug(
			"ARP RESULT:",
			ip,
			mac,
		)

		if mac != "" {

			host.MAC = mac

			host.Online = true

			host.Vendor = LookupVendor(mac)

			host.Sources = append(
				host.Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryARP,
					Value: mac,
				},
			)

			if host.Vendor != "" && host.Vendor != "Unknown" {

				host.Sources = append(
					host.Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryOUI,
						Value: host.Vendor,
					},
				)
			}
		}
	}

	// Reverse DNS

	if config.EnableReverseDNS {

		stepStart := time.Now()

		hostname := LookupReverseDNS(ip)

		dnsDuration := time.Since(stepStart)

		if dnsDuration > 100*time.Millisecond {

			logger.Debug(
				"DISCOVERY SLOW:",
				ip,
				"REVERSE DNS:",
				dnsDuration,
			)
		}

		if hostname != "" && host.Hostname == "" {

			host.Hostname = hostname

			host.Sources = append(
				host.Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryReverseDNS,
					Value: hostname,
				},
			)
		}
	}
	totalDuration := time.Since(start)

	if totalDuration > 1*time.Second {

		logger.Debug(
			"DISCOVERY SLOW TOTAL:",
			ip,
			totalDuration,
		)
	}

	return host

}
