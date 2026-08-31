package scanner

import (
	"sync"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func udpDiscoveryWorker(
	jobs <-chan UDPDiscoveryJob,
	results chan<- UDPDiscoveryResult,
	iface models.Interface,
	config models.ScannerConfig,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for job := range jobs {

		services := DiscoverUDP(
			job.IP,
			iface,
			config,
		)

		results <- UDPDiscoveryResult{
			Index:    job.Index,
			Services: services,
		}
	}
}

func DiscoverHostsFull(
	ips []string,
	workers int,
	config models.ScannerConfig,
) []models.Host {

	if workers <= 0 {
		workers = 1
	}

	logger.Info(
		"DISCOVERY WORKERS:",
		workers,
	)

	jobs := make(chan string)
	results := make(chan models.Host)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go discoveryWorker(
			jobs,
			results,
			&wg,
			config,
		)
	}

	go func() {

		for _, ip := range ips {
			jobs <- ip
		}

		close(jobs)

	}()

	go func() {

		wg.Wait()
		close(results)

	}()

	var hosts []models.Host

	for host := range results {

		logger.Debug(
			"DISCOVERY RESULT:",
			host.IP,
			"ONLINE:",
			host.Online,
			"MAC:",
			host.MAC,
		)

		if !IsHostDiscovered(host) {
			continue
		}

		hosts = append(hosts, host)
	}

	return hosts
}

func IsHostDiscovered(
	host models.Host,
) bool {

	if host.Online {
		return true
	}

	if host.MAC != "" {
		return true
	}

	if len(host.Ports) > 0 {
		return true
	}

	if len(host.UDPServices) > 0 {
		return true
	}

	if len(host.HTTP) > 0 {
		return true
	}

	return false
}
