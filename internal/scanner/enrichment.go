package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

// EnrichHosts collects additional information
// for every discovered host.
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


    // ==========================
    // Final host state
    // ==========================

    for i := range hosts {

        hosts[i].Confidence =
            CalculateConfidence(
                hosts[i],
            )

        logger.Log.Println(
            "HOST FINAL:",
            hosts[i].IP,
            hosts[i].MAC,
            hosts[i].Vendor,
            hosts[i].Hostname,
            hosts[i].Type,
            "confidence:",
            hosts[i].Confidence,
        )
    }


    return hosts
}