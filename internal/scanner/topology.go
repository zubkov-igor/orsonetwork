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

import "OrsoNetwork/internal/logger"
import "OrsoNetwork/internal/models"

func BuildTopology(
	networks []models.Network,
) models.Topology {

	topology := models.Topology{
		Networks: networks,
	}

	for _, network := range networks {

		var gateway models.Host

		// Находим объект шлюза.
		for _, h := range network.Hosts {

			if h.IP == network.Gateway {
				gateway = h
				break
			}
		}

		for _, host := range network.Hosts {

			nodeType := host.Type

			if host.IP == network.Gateway {
				nodeType = models.DeviceGateway
			}

			label := host.IP

			if host.Hostname != "" {
				label = host.Hostname
			}

			logger.Debug(
				"TOPOLOGY HOST STATUS:",
				host.IP,
				"ONLINE:",
				host.Online,
				"TYPE:",
				host.Type,
			)

			topology.Nodes = append(
				topology.Nodes,
				models.Node{
					ID:       NodeID(host),
					Label:    label,
					Type:     string(nodeType),
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

	for _, node := range topology.Nodes {

		logger.Debug(
			"TOPOLOGY FINAL NODE:",
			"ID:", node.ID,
			"IP:", node.IP,
			"ONLINE:", node.Online,
			"TYPE:", node.Type,
		)
	}

	return topology
}

func NodeID(host models.Host) string {

	if host.MAC != "" {
		return host.MAC
	}

	return host.IP
}
