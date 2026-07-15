package topology

func BuildTopology(
    networks []models.Network,
) models.Topology {

    topology := models.Topology{
        Networks: networks,
    }

    for _, network := range networks {

        for _, host := range network.Hosts {

            if host.IP == network.Gateway {
                continue
            }

            topology.Links = append(
                topology.Links,
                models.Link{
                    From: network.Gateway,
                    To:   host.IP,
                    Type: "network",
                },
            )
        }
    }

    return topology
}