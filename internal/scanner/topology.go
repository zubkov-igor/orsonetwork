// BuildTopology converts discovered hosts
// into a graph representation.
//
// Current links describe logical network relation
// through the gateway.
//
// TODO:
// improve topology discovery using:
// - LLDP
// - SNMP
// - ARP relationships
// - WiFi information

package scanner

import "OrsoNetwork/internal/models"

func BuildTopology(
	networks []models.Network,
) models.Topology {

	topology := models.Topology{
		Networks: networks,
	}

	for _, network := range networks {

		var gateway models.Host

		// находим объект шлюза
		for _, h := range network.Hosts {

			if h.IP == network.Gateway {
				gateway = h
				break
			}
		}


        for _, host := range network.Hosts {

            label := host.IP

            if host.Hostname != "" {
                label = host.Hostname
            }

            topology.Nodes = append(
                topology.Nodes,
                models.Node{
                    ID:       NodeID(host),
                    Label:    label,
                    Type:     string(host.Type),
                    IP:       host.IP,
                    MAC:      host.MAC,
                    Hostname: host.Hostname,
                    Vendor:   host.Vendor,
                    Sources:  host.Sources,
                    Online:   host.Online,
                    RTT:      host.RTT,
                },
            )


            if host.IP != network.Gateway && gateway.IP != "" {

                topology.Links = append(
                    topology.Links,
                    models.Link{
                        From: NodeID(gateway),
                        To:   NodeID(host),
                        Type: "network",
                    },
                )
            }
        }
    }


    return topology
}


func NodeID(host models.Host) string {

	if host.MAC != "" {
		return host.MAC
	}

	return host.IP
}