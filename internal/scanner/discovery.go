package scanner

import (
    "sync"
    "time"

    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

const pingWorkers = 24
const pingTimeout = 200 * time.Millisecond


func DiscoverHosts(
    cidr string,
    ownIP string,
) []models.Host {

    hosts := []models.Host{}

    start := time.Now()

    ips := HostsFromCIDR(cidr)

    logger.Log.Println(
        "CIDR HOSTS GENERATED:",
        len(ips),
        time.Since(start),
    )


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


    jobs := make(chan string)
    results := make(chan models.Host)


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

        hosts = append(
            hosts,
            host,
        )
    }


    logger.Log.Println(
        "PING DISCOVERY FINISHED:",
        len(hosts),
        time.Since(start),
    )


    return hosts
}