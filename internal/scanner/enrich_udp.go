package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichUDP(
	hosts []models.Host,
	iface models.Interface,
) []models.Host {

	logger.Log.Println(
		"UDP ENRICHMENT START",
	)

	for i := range hosts {

		services := DiscoverUDP(
			hosts[i].IP,
			iface,
		)

		hosts[i].UDPServices = services

		for _, u := range services {

			logger.Log.Println(
				"UDP SERVICE:",
				hosts[i].IP,
				u.Port,
				u.Service,
			)

			hosts[i].Sources = append(
				hosts[i].Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryUDP,
					Value: u.Service,
				},
			)
		}
	}

	logger.Log.Println(
		"UDP ENRICHMENT FINISHED",
	)

	return hosts
}
