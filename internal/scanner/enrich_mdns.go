package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichMDNS(
	hosts []models.Host,
) []models.Host {

	logger.Debug(
		"MDNS ENRICHMENT START",
	)

	mdnsRecords := DiscoverMDNS()

	logger.Info(
		"MDNS FOUND:",
		len(mdnsRecords),
	)

	for i := range hosts {

		for _, mdns := range mdnsRecords {

			logger.Debug(
				"MDNS COMPARE:",
				"host=", hosts[i].IP,
				"mdns=", mdns.IP,
			)

			if mdns.IP != hosts[i].IP {
				continue
			}

			hosts[i].MDNS = append(
				hosts[i].MDNS,
				mdns,
			)

			logger.Info(
				"MDNS MATCH:",
				hosts[i].IP,
				mdns.Name,
				mdns.Service,
				mdns.Host,
				mdns.Port,
			)
		}
	}

	logger.Info(
		"MDNS ENRICHMENT FINISHED",
	)

	return hosts
}
