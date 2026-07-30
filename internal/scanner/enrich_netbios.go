package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)


func EnrichNetBIOS(
    hosts []models.Host,
) []models.Host {


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


        if netbios.Name != "" {


            if hosts[i].Hostname == "" {

                hosts[i].Hostname =
                    netbios.Name
            }


            hosts[i].Sources = append(
                hosts[i].Sources,
                models.DiscoverySource{
                    Type: models.DiscoveryNetBIOS,
                    Value: netbios.Name,
                },
            )
        }


        // fallback MAC
        // если ARP не дал MAC,
        // но NetBIOS дал

        if hosts[i].MAC == "" &&
            netbios.MAC != "" {


            hosts[i].MAC =
                netbios.MAC


            hosts[i].Vendor =
                LookupVendor(
                    netbios.MAC,
                )


            hosts[i].Sources = append(
                hosts[i].Sources,
                models.DiscoverySource{
                    Type: models.DiscoveryNetBIOS,
                    Value: netbios.MAC,
                },
            )
        }
    }


    return hosts
}