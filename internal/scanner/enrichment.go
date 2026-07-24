package scanner

import "OrsoNetwork/internal/models"

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
    
}

	return hosts
}
