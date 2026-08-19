package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

// EnrichARP adds MAC address and vendor information
// using ARP discovery.
//
// Discovery sources:
// ARP  -> MAC address
// OUI  -> Vendor

func EnrichARP(
	hosts []models.Host,
) []models.Host {

	logger.Log.Println(
		"ARP ENRICHMENT START",
	)

	arpHosts := ARPDiscovery(hosts)

	arpMap := make(
		map[string]models.Host,
	)

	for _, h := range arpHosts {
		arpMap[h.IP] = h
	}

	for i := range hosts {

		if arpHost, ok := arpMap[hosts[i].IP]; ok {

			if arpHost.MAC == "" {
				continue
			}

			hosts[i].MAC = arpHost.MAC

			hosts[i].Sources = append(
				hosts[i].Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryARP,
					Value: arpHost.MAC,
				},
			)

			hosts[i].Vendor = LookupVendor(
				arpHost.MAC,
			)

			if hosts[i].Vendor != "" {

				hosts[i].Sources = append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryOUI,
						Value: hosts[i].Vendor,
					},
				)
			}

			logger.Log.Println(
				"ARP ENRICHED:",
				hosts[i].IP,
				hosts[i].MAC,
				hosts[i].Vendor,
			)
		}
	}

	logger.Log.Println(
		"ARP ENRICHMENT FINISHED",
	)

	return hosts
}
