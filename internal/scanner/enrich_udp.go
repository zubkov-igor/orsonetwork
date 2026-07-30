package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func EnrichUDP(
    hosts []models.Host,
) []models.Host {


    for i := range hosts {


        hosts[i].UDPServices = DiscoverUDP(
            hosts[i].IP,
        )


        for _, u := range hosts[i].UDPServices {

            logger.Log.Println(
                "UDP SERVICE:",
                hosts[i].IP,
                u.Port,
                u.Service,
            )
        }


        if len(hosts[i].UDPServices) > 0 {

            hosts[i].Sources = append(
                hosts[i].Sources,
                models.DiscoverySource{
                    Type: models.DiscoveryUDP,
                    Value: "UDP discovery",
                },
            )
        }
    }


    return hosts
}