package scanner

import (
	"fmt"
	"net"
	"net/netip"
	"sync"
	"time"

	"OrsoNetwork/internal/models"

	"github.com/mdlayher/arp"
)

func arpWorker(
	client *arp.Client,
	jobs <-chan string,
	results chan<- models.Host,
	wg *sync.WaitGroup,
) {

	defer wg.Done()
	defer client.Close()

	for ip := range jobs {

		addr, err := netip.ParseAddr(ip)

		if err != nil {
			continue
		}

		mac, err := client.Resolve(addr)

		if err != nil {
			continue
		}

		results <- models.Host{
			IP:     ip,
			MAC:    mac.String(),
			Online: true,
		}

	}

}

func arpResolve(
	client *arp.Client,
	addr netip.Addr,
) (net.HardwareAddr, error) {

	type result struct {
		mac net.HardwareAddr
		err error
	}

	ch := make(chan result, 1)

	go func() {

		mac, err := arpResolve(
			client,
			addr,
		)

		ch <- result{
			mac: mac,
			err: err,
		}

	}()

	select {

	case r := <-ch:
		return r.mac, r.err

	case <-time.After(
		500 * time.Millisecond,
	):
		return nil, fmt.Errorf("arp timeout")

	}

}
