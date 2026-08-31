package scanner

import (
	"fmt"
	"sync"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

type portScanResult struct {
	index int
	ports []models.Port
}

func EnrichPorts(
	hosts []models.Host,
) []models.Host {

	logger.Info(
		"PORT ENRICHMENT START",
	)

	const workers = 8

	type portScanResult struct {
		index int
		ports []models.Port
	}

	jobs := make(chan int)
	results := make(chan portScanResult)

	var wg sync.WaitGroup

	// Start workers.

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for index := range jobs {

				logger.Debug(
					"PORT SCAN:",
					hosts[index].IP,
				)

				ports := ScanPorts(
					hosts[index].IP,
				)

				results <- portScanResult{
					index: index,
					ports: ports,
				}
			}

		}()
	}

	// Send hosts to workers.

	go func() {

		for i := range hosts {

			jobs <- i
		}

		close(jobs)

	}()

	// Close results after workers finish.

	go func() {

		wg.Wait()
		close(results)

	}()

	// Process results.

	for result := range results {

		i := result.index
		ports := result.ports

		hosts[i].Ports = ports

		for _, p := range ports {

			logger.Info(
				"OPEN PORT:",
				hosts[i].IP,
				p.Number,
				p.Protocol,
				p.Service,
			)

			hosts[i].Sources = append(
				hosts[i].Sources,
				models.DiscoverySource{
					Type: models.DiscoveryTCP,
					Value: fmt.Sprintf(
						"%s:%d:%s",
						p.Protocol,
						p.Number,
						p.Service,
					),
				},
			)

			if p.Service == "http" {

				httpInfo := ScanHTTP(
					hosts[i].IP,
					p.Number,
				)

				if httpInfo.Server != "" ||
					httpInfo.Title != "" ||
					len(httpInfo.Scripts) > 0 ||
					len(httpInfo.Keywords) > 0 {

					hosts[i].HTTP = append(
						hosts[i].HTTP,
						httpInfo,
					)

					logger.Info(
						"HTTP ENRICHED:",
						hosts[i].IP,
						httpInfo.Port,
						httpInfo.Server,
						httpInfo.Title,
						httpInfo.Keywords,
					)
				}
			}
		}
	}

	logger.Info(
		"PORT ENRICHMENT FINISHED",
	)

	return hosts
}
