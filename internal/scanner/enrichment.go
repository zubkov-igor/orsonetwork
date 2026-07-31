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


    logger.Section("ARP ENRICHMENT")
    hosts = EnrichARP(hosts)


    logger.Section("NETBIOS ENRICHMENT")
    hosts = EnrichNetBIOS(hosts)


    logger.Section("MDNS DISCOVERY")
    hosts = EnrichMDNS(hosts)


    logger.Section("UDP DISCOVERY")
    hosts = EnrichUDP(hosts)


    logger.Section("REVERSE DNS")
    hosts = EnrichReverseDNS(hosts)


    logger.Section("TCP PORT SCAN")
    hosts = EnrichPorts(hosts)



    logger.Section("CONFIDENCE CALCULATION")


    for i := range hosts {

        hosts[i].Type =
            IdentifyDevice(
        hosts[i],
    )

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