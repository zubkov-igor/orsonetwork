package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func ScanPorts(
	ip string,
) []models.Port {

	scanStart := time.Now()

	logger.Debug(
		"PORT SCAN START:",
		ip,
	)

	const workers = 18

	jobs := make(chan int)
	results := make(chan models.Port)

	var wg sync.WaitGroup

	// Start workers.

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go scanPortWorker(
			ip,
			jobs,
			results,
			&wg,
		)
	}

	// Send ports to workers.

	go func() {

		for _, port := range CommonPorts {
			jobs <- port
		}

		close(jobs)

	}()

	// Close results after all workers finish.

	go func() {

		wg.Wait()
		close(results)

	}()

	var ports []models.Port

	for port := range results {

		ports = append(
			ports,
			port,
		)
	}

	logger.Debug(
		"PORT SCAN FINISHED:",
		ip,
		len(ports),
		"DURATION:",
		time.Since(scanStart),
	)

	return ports
}

func scanPortWorker(
	ip string,
	jobs <-chan int,
	results chan<- models.Port,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for port := range jobs {

		if !IsPortOpen(
			ip,
			port,
		) {
			continue
		}

		banner := GrabBanner(
			ip,
			port,
		)

		logger.Debug(
			"PORT BANNER:",
			ip,
			port,
			banner,
		)

		results <- models.Port{
			Number:   port,
			Protocol: "tcp",
			Service:  DetectPortService(port),
			Open:     true,
			Banner:   banner,
		}
	}
}

func IsPortOpen(
	ip string,
	port int,
) bool {

	address := fmt.Sprintf(
		"%s:%d",
		ip,
		port,
	)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		500*time.Millisecond,
	)

	if err != nil {
		return false
	}

	conn.Close()

	return true
}
