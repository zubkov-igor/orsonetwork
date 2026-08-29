package scanner

import (
	"sync"

    "OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func discoveryWorker(
    jobs <-chan string,
    results chan<- models.Host,
    wg *sync.WaitGroup,
    config models.ScannerConfig,
) {

    defer wg.Done()

    for ip := range jobs {

        host := discoverHostNew(
            ip,
            config,
        )

        results <- host
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

    logger.Log.Println(
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

        hosts = append(
            hosts,
            host,
        )
    }

    return hosts
}