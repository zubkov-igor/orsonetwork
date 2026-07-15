package scanner

import "OrsoNetwork/internal/models"

func EnrichHosts(hosts []models.Host) []models.Host {

	for i := range hosts {

		hosts[i].Vendor = LookupVendor(
			hosts[i].MAC,
		)

		hosts[i].Hostname = LookupHostname(
			hosts[i].IP,
		)

	}

	return hosts
}