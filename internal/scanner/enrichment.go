package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func EnrichHosts(
	hosts []models.Host,
) []models.Host {

	arpHosts := ARPDiscovery(
		hosts,
	)

	arpMap := make(
		map[string]models.Host,
	)

	for _, h := range arpHosts {
		arpMap[h.IP] = h
	}

for i := range hosts {

    // ARP
    if arpHost, ok := arpMap[hosts[i].IP]; ok {

        if arpHost.MAC != "" {

            hosts[i].MAC = arpHost.MAC

            hosts[i].Sources = append(
                hosts[i].Sources,
                models.DiscoverySource{
                    Type:  models.DiscoveryARP,
                    Value: arpHost.MAC,
                },
            )

            hosts[i].Vendor = LookupVendor(
                arpHost.MAC,
            )
        }
    }


    // Reverse DNS
    hostname := LookupReverseDNS(
        hosts[i].IP,
    )

    if hostname != "" {

        hosts[i].Hostname = hostname

        hosts[i].Sources = append(
            hosts[i].Sources,
            models.DiscoverySource{
                Type:  models.DiscoveryReverseDNS,
                Value: hostname,
            },
        )
    }


    // NetBIOS fallback
    if hosts[i].Hostname == "" {

        netbios, err := LookupNetBIOS(
            hosts[i].IP,
        )

        if err == nil {

            if netbios.Name != "" {

                hosts[i].Hostname = netbios.Name

                hosts[i].Sources = append(
                    hosts[i].Sources,
                    models.DiscoverySource{
                        Type: models.DiscoveryNetBIOS,
                        Value: netbios.Name,
                    },
                )
            }


            if hosts[i].MAC == "" && netbios.MAC != "" {

                hosts[i].MAC = netbios.MAC

                hosts[i].Sources = append(
                    hosts[i].Sources,
                    models.DiscoverySource{
                        Type: models.DiscoveryNetBIOS,
                        Value: netbios.MAC,
                    },
                )
            }
        }
    }

hosts[i].Ports = ScanPorts(
    hosts[i].IP,
)

for _, p := range hosts[i].Ports {

    logger.Log.Println(
        "OPEN PORT:",
        hosts[i].IP,
        p.Number,
        p.Protocol,
        p.Service,
    )
}

// Device detection
device := models.Device{
    IP:       hosts[i].IP,
    MAC:      hosts[i].MAC,
    Hostname: hosts[i].Hostname,
}

hosts[i].Type = IdentifyDevice(device)

logger.Log.Println(
    "DEVICE DETECTED:",
    hosts[i].IP,
    hosts[i].Hostname,
    hosts[i].Vendor,
    hosts[i].Type,
)

}

	return hosts
}
