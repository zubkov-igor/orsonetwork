package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func EnrichMDNS(
    hosts []models.Host,
) []models.Host {


    logger.Log.Println(
        "MDNS ENRICHMENT START",
    )



    mdnsRecords := DiscoverMDNS()


    logger.Log.Println(
        "MDNS FOUND:",
        len(mdnsRecords),
    )



    for i := range hosts {


        for _, mdns := range mdnsRecords {


            if mdns.IP != hosts[i].IP {
                continue
            }



            logger.Log.Println(
                "MDNS MATCH:",
                hosts[i].IP,
                mdns.Name,
            )



            hosts[i].MDNS =
                append(
                    hosts[i].MDNS,
                    mdns,
                )



            if mdns.Name != "" {

                hosts[i].Sources =
                    append(
                        hosts[i].Sources,
                        models.DiscoverySource{
                            Type:  models.DiscoveryMDNS,
                            Value: mdns.Name,
                        },
                    )
            }
        }
    }



    logger.Log.Println(
        "MDNS ENRICHMENT FINISHED",
    )


    return hosts
}