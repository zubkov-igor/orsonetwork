package scanner

import (
    "sync"
    "time"
    "fmt"

    "OrsoNetwork/internal/models"
)


const workers = 24

const timeout = 200 * time.Millisecond


func DiscoverHosts(cidr string) []models.Host {


    var hosts []models.Host

    start := time.Now()

    ips := HostsFromCIDR(cidr)

    fmt.Println("CIDR:", time.Since(start))

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

for _, h := range hosts {

    fmt.Println(
        "ENRICH RESULT:",
        h.IP,
        h.MAC,
        h.Vendor,
    )
}

    for host := range results {

        hosts = append(
            hosts,
            host,
        )

    }

    fmt.Println("Ping:", time.Since(start))


    return hosts
}