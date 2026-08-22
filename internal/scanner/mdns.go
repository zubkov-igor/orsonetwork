package scanner

import (
	"context"
	"sync"
	"time"

	"github.com/grandcat/zeroconf"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func DiscoverMDNS() []models.MDNSService {

	var records []models.MDNSService

	resolver, err := zeroconf.NewResolver(nil)

	if err != nil {
		return records
	}

	entries := make(chan *zeroconf.ServiceEntry)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		for entry := range entries {

			logger.Log.Println(
				"MDNS SERVICE:",
				entry.Instance,
				entry.Service,
				entry.HostName,
				entry.Port,
				entry.AddrIPv4,
			)

			for _, ip := range entry.AddrIPv4 {

				records = append(
					records,
					models.MDNSService{
						Name:    entry.Instance,
						Service: entry.Service,
						Host:    entry.HostName,
						IP:      ip.String(),
						Port:    entry.Port,
					},
				)
			}
		}
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	_ = resolver.Browse(
		ctx,
		"_dosvc._tcp",
		"local",
		entries,
	)

	wg.Wait()

	return records
}
