package scanner

import "OrsoNetwork/internal/models"

func EnrichHosts(
	hosts []models.Host,
	cidr string,
) []models.Host {

	arpHosts := ARPDiscovery(cidr)

	arpMap := make(
		map[string]models.Host,
	)

	for _, h := range arpHosts {
		arpMap[h.IP] = h
	}

	for i := range hosts {

		if arpHost, ok := arpMap[hosts[i].IP]; ok {

			hosts[i].MAC = arpHost.MAC

			hosts[i].Vendor = LookupVendor(
				arpHost.MAC,
			)
		}
	}

	return hosts
}