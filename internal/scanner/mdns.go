package scanner

import (
	"context"

	"github.com/grandcat/zeroconf"

	"OrsoNetwork/internal/models"
)

func DiscoverMDNS() []models.MDNSService {

	var records []models.MDNSService

	resolver, err := zeroconf.NewResolver(nil)

	if err != nil {
		return records
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func() {

		for entry := range entries {

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
		5,
	)

	defer cancel()

	_ = resolver.Browse(
		ctx,
		"_services._dns-sd._udp",
		"local",
		entries,
	)

	return records
}
