package scanner

import (
	"sort"
	"sync"
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

// Number of concurrent ping workers.
//
// Controls scan speed and CPU/network load.

const pingWorkers = 24
const pingTimeout = 200 * time.Millisecond

// DiscoverHosts scans a network using ICMP echo requests.
//
// Pipeline:
//
// CIDR
//     ↓
// Generate IPs
//     ↓
// Remove own IP
//     ↓
// Worker pool
//     ↓
// Alive hosts
//
// Returns only hosts that responded.

func DiscoverHosts(
	cidr string,
	ownIP string,
) []models.Host {

	hosts := []models.Host{}

	start := time.Now()

	// Generate every host address
	// inside the network.

	ips := HostsFromCIDR(cidr)

	logger.Log.Println(
		"CIDR HOSTS GENERATED:",
		len(ips),
		time.Since(start),
	)

	// Never ping ourselves.

	filteredIPs := make([]string, 0, len(ips))

	for _, ip := range ips {

		if ip == ownIP {

			logger.Log.Println(
				"SKIP OWN IP:",
				ip,
			)

			continue
		}

		filteredIPs = append(
			filteredIPs,
			ip,
		)
	}

	ips = filteredIPs

	logger.Log.Println(
		"CIDR HOSTS AFTER FILTER:",
		len(ips),
		time.Since(start),
	)

	// jobs
	// IPs waiting to be scanned.
	//
	// results
	// Alive hosts discovered by workers.

	jobs := make(chan string)
	results := make(chan models.Host)

	// Wait until every worker
	// finishes processing.

	var wg sync.WaitGroup

	wg.Add(pingWorkers)

	for i := 0; i < pingWorkers; i++ {

		go pingWorker(
			jobs,
			results,
			&wg,
			pingTimeout,
		)
	}

	// Feed every IP into the worker pool.

	go func() {

		for _, ip := range ips {
			jobs <- ip
		}

		close(jobs)

	}()

	// Close result channel
	// after all workers finish.

	go func() {

		wg.Wait()

		close(results)

	}()

	// Collect alive hosts.

	for host := range results {

		hosts = append(
			hosts,
			host,
		)
	}

	sort.Slice(hosts, func(i, j int) bool {
		return hosts[i].IP < hosts[j].IP
	})

	logger.Log.Println(
		"PING DISCOVERY FINISHED:",
		len(hosts),
		time.Since(start),
	)

	return hosts
}
