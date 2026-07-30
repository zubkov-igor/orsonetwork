package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func EnrichMDNS(
    hosts []models.Host,
) []models.Host {


    mdnsRecords := DiscoverMDNS()


    logger.Log.Println(
        "MDNS FOUND:",
        len(mdnsRecords),
    )


    for i := range hosts {

        for _, mdns := range mdnsRecords {

            if mdns.IP == hosts[i].IP {


                hosts[i].MDNS = append(
                    hosts[i].MDNS,
                    mdns,
                )


                hosts[i].Sources = append(
                    hosts[i].Sources,
                    models.DiscoverySource{
                        Type: models.DiscoveryMDNS,
                        Value: mdns.Name,
                    },
                )
            }
        }
    }


    return hosts
}