package scanner

import (
	"sync"
	"time"

	"OrsoNetwork/internal/models"
)


func worker(
    jobs <-chan string,
    results chan<- models.Host,
    wg *sync.WaitGroup,
    timeout time.Duration,
) {

	defer wg.Done()


	for ip := range jobs {

		host := PingHost(
			ip,
			timeout,
		)


		if host.Online {

			results <- host

		}
	}
}