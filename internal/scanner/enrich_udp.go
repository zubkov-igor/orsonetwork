package scanner

import (
	"sync"
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichUDP(
	hosts []models.Host,
	iface models.Interface,
) []models.Host {

	logger.Log.Println(
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

	// Start workers.

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go udpDiscoveryWorker(
			jobs,
			results,
			iface,
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

			logger.Log.Println(
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

	// mDNS remains a separate discovery step
	// for now, even though it is currently called
	// from this enrichment stage.

	mdnsHosts := ProbeMDNS(
		iface,
	)

	logger.Log.Println(
		"MDNS FOUND:",
		len(mdnsHosts),
	)

	for _, ip := range mdnsHosts {

		logger.Log.Println(
			"MDNS HOST:",
			ip,
		)
	}

	logger.Log.Println(
		"UDP ENRICHMENT FINISHED",
	)

	return hosts
}