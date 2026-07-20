package scanner

import (
	"fmt"
	"sync"
	"time"

	"OrsoNetwork/internal/models"
)

const workers = 24

const timeout = 200 * time.Millisecond

func DiscoverHosts(
	cidr string,
	ownIP string,
) []models.Host {

	var hosts []models.Host

	start := time.Now()

	ips := HostsFromCIDR(cidr)

	fmt.Println(
		"CIDR:",
		time.Since(start),
	)

	jobs := make(chan string)

	results := make(chan models.Host)

	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go worker(
			jobs,
			results,
			&wg,
			timeout,
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

	for host := range results {

		if host.IP == ownIP {

			fmt.Println(
				"SKIP OWN HOST:",
				host.IP,
			)

			continue
		}

		hosts = append(
			hosts,
			host,
		)
	}

	fmt.Println(
		"Ping:",
		time.Since(start),
	)

	return hosts
}
