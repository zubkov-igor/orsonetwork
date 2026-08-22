package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichHosts(
	hosts []models.Host,
	iface models.Interface,
) []models.Host {

	hosts = EnrichARP(
		hosts,
	)

	hosts = EnrichNetBIOS(
    hosts,
)

logger.Log.Println(
    "HOSTS BEFORE MDNS:",
    len(hosts),
)

for _, host := range hosts {

    logger.Log.Println(
        "HOST BEFORE MDNS:",
        host.IP,
    )
}

hosts = EnrichMDNS(
    hosts,
)

	hosts = EnrichUDP(
		hosts,
		iface,
	)

	hosts = EnrichReverseDNS(
		hosts,
	)

	hosts = EnrichPorts(
		hosts,
	)

	for i := range hosts {

		hosts[i].Type = IdentifyDevice(
			hosts[i],
		)

		hosts[i].Confidence =
			CalculateConfidence(
				hosts[i],
			)
	}

	return hosts
}
