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

            nodeType := "host"

            if host.IP == network.Gateway {
                nodeType = "gateway"
            }

            topology.Nodes = append(
                topology.Nodes,
                models.Node{
                    ID:     NodeID(host),
                    Label:  host.IP,
                    Type:   nodeType,
                    IP:     host.IP,
                    MAC:    host.MAC,
                    Vendor: host.Vendor,
                },
            )

            if host.IP != network.Gateway {

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