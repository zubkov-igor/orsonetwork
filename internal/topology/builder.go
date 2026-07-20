package topology

import (
	"OrsoNetwork/internal/models"
)

func BuildTopology(
	networks []models.Network,
) models.Topology {

	topology := models.Topology{
		Networks: networks,
	}

	for _, network := range networks {

		for _, host := range network.Hosts {

			nodeType := "host"

			if host.IP == network.Gateway {
				nodeType = "gateway"
			}

			topology.Nodes = append(
				topology.Nodes,
				models.Node{
					ID:    NodeID(host),
					Label: host.IP,
					Type:  nodeType,

					IP:       host.IP,
					MAC:      host.MAC,
					Hostname: host.Hostname,
					Vendor:   host.Vendor,

					Sources: host.Sources,

					Online: host.Online,
					RTT:    host.RTT,
				},
			)

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
