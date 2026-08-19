package scanner

import (
	"OrsoNetwork/internal/models"
)

func EnrichReverseDNS(
	hosts []models.Host,
) []models.Host {

	for i := range hosts {

		hostname := LookupReverseDNS(
			hosts[i].IP,
		)

		if hostname == "" {
			continue
		}

		hosts[i].Hostname = hostname

		hosts[i].Sources = append(
			hosts[i].Sources,
			models.DiscoverySource{
				Type:  models.DiscoveryReverseDNS,
				Value: hostname,
			},
		)
	}

	return hosts
}
