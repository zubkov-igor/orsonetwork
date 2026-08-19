package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichNetBIOS(
	hosts []models.Host,
) []models.Host {

	logger.Log.Println(
		"NETBIOS ENRICHMENT START",
	)

	for i := range hosts {

		netbios, err := LookupNetBIOS(
			hosts[i].IP,
		)

		if err != nil {
			continue
		}

		logger.Log.Println(
			"NETBIOS FOUND:",
			hosts[i].IP,
			netbios.Name,
			netbios.MAC,
		)

		// Hostname enrichment

		if netbios.Name != "" {

			if hosts[i].Hostname == "" {

				hosts[i].Hostname =
					netbios.Name
			}

			hosts[i].Sources =
				append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryNetBIOS,
						Value: netbios.Name,
					},
				)
		}

		// MAC fallback
		//
		// ARP has priority.
		// NetBIOS fills missing MAC.

		if hosts[i].MAC == "" &&
			netbios.MAC != "" {

			hosts[i].MAC =
				netbios.MAC

			hosts[i].Sources =
				append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryNetBIOS,
						Value: netbios.MAC,
					},
				)

			// Vendor lookup from MAC

			hosts[i].Vendor =
				LookupVendor(
					netbios.MAC,
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
		}
	}

	logger.Log.Println(
		"NETBIOS ENRICHMENT FINISHED",
	)

	return hosts
}
