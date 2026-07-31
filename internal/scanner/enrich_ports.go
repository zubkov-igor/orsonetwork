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

			if p.Service == "http" {

				httpInfo := ScanHTTP(
					hosts[i].IP,
					p.Number,
				)

				if httpInfo.Server != "" ||
					httpInfo.Title != "" ||
					len(httpInfo.Scripts) > 0 ||
					len(httpInfo.Keywords) > 0 {

					hosts[i].HTTP = append(
						hosts[i].HTTP,
						httpInfo,
					)

					logger.Log.Println(
						"HTTP ENRICHED:",
						hosts[i].IP,
						httpInfo.Port,
						httpInfo.Server,
						httpInfo.Title,
						httpInfo.Keywords,
					)
				}
			}
		}

		// здесь заканчивается цикл по портам

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