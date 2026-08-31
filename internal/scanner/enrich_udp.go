package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
	"sync"
)

func EnrichUDP(
	hosts []models.Host,
	iface models.Interface,
	config models.ScannerConfig,
) []models.Host {

	logger.Info(
		"UDP ENRICHMENT START",
	)

	workers := 20

	jobs := make(
		chan UDPDiscoveryJob,
	)

	results := make(
		chan UDPDiscoveryResult,
	)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go udpDiscoveryWorker(
			jobs,
			results,
			iface,
			config,
			&wg,
		)
	}

	// Send jobs.

	go func() {

		for i, host := range hosts {

			jobs <- UDPDiscoveryJob{
				Index: i,
				IP:    host.IP,
			}
		}

		close(jobs)
	}()

	// Close results after workers finish.

	go func() {

		wg.Wait()

		close(results)
	}()

	// Collect results.

	for result := range results {

		i := result.Index

		hosts[i].UDPServices =
			result.Services

		for _, u := range result.Services {

			logger.Info(
				"UDP SERVICE:",
				hosts[i].IP,
				u.Port,
				u.Service,
			)

			hosts[i].Sources =
				append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryUDP,
						Value: u.Service,
					},
				)
		}
	}

	logger.Info(
		"UDP ENRICHMENT FINISHED",
	)

	return hosts
}

func EnrichSNMP(
	hosts []models.Host,
) []models.Host {

	logger.Info(
		"SNMP ENRICHMENT START",
	)

	for i := range hosts {

		result := ProbeSNMP(
			hosts[i].IP,
		)

		if result.Found {

			hosts[i].UDPServices =
				append(
					hosts[i].UDPServices,
					models.UDPService{
						IP:       hosts[i].IP,
						Port:     161,
						Service:  "SNMP",
						Protocol: "udp",
						Info:     result.Info,
					},
				)

			hosts[i].Sources =
				append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoverySNMP,
						Value: "SNMP",
					},
				)
		}
	}

	logger.Info(
		"SNMP ENRICHMENT FINISHED",
	)

	return hosts
}
