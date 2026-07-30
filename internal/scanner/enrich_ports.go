package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichPorts(
	hosts []models.Host,
) []models.Host {

	for i := range hosts {

		ports := ScanPorts(
			hosts[i].IP,
		)

		hosts[i].Ports = ports

		for _, p := range ports {

			logger.Log.Println(
				"OPEN PORT:",
				hosts[i].IP,
				p.Number,
				p.Protocol,
				p.Service,
			)
		}

		if len(ports) > 0 {

			hosts[i].Sources = append(
				hosts[i].Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryTCP,
					Value: "TCP ports",
				},
			)
		}
	}

	return hosts
}