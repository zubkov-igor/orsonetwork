package scanner

import (
    "OrsoNetwork/internal/models"
)

func EnrichHosts(
    hosts []models.Host,
) []models.Host {


    hosts = EnrichARP(
        hosts,
    )


    hosts = EnrichNetBIOS(
        hosts,
    )


    hosts = EnrichMDNS(
        hosts,
    )


    hosts = EnrichUDP(
        hosts,
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