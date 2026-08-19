package scanner

import (
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

// Scanner is the main orchestrator.
//
// It coordinates the entire discovery process
// but does not implement discovery itself.
//
// Every step is delegated to a dedicated module.

type Scanner struct {
}

func New() *Scanner {
	return &Scanner{}
}

// Scan performs full network discovery.
//
// Pipeline:
//
// Interfaces
//     ↓
// Gateways
//     ↓
// Networks
//     ↓
// Host discovery
//     ↓
// Host enrichment
//
// Returns all discovered networks.

func (s *Scanner) Scan() []models.Network {

	logger.Separator(
		"NEW SCAN START",
	)

	logger.Section(
		"NETWORK DISCOVERY",
	)

	var networks []models.Network

	// Discover available network interfaces.

	interfaces := GetInterfaces()

	logger.Log.Println(
		"SCANNER INTERFACES:",
		len(interfaces),
	)

	gateways := GetGateways()

	logger.Log.Println(
		"SCANNER GATEWAYS:",
		len(gateways),
	)

	logger.Log.Println(
		"SCANNER LOOP START",
	)

	for _, iface := range interfaces {

		logger.Log.Println(
			"PROCESS INTERFACE:",
			iface.Name,
		)

		gw := GatewayForInterface(
			iface,
			gateways,
		)

		if gw == nil {

			logger.Log.Println(
				"NO GATEWAY FOR INTERFACE:",
				iface.Name,
			)

			continue
		}

		logger.Log.Println(
			"GATEWAY FOUND:",
			iface.Name,
			gw.IP,
		)

		network := BuildNetwork(
			iface,
			gw,
		)

		logger.Log.Println(
			"NETWORK BUILT:",
			network.CIDR,
		)

		network.Hosts = DiscoverHosts(
			network.CIDR,
			iface.IP,
		)

		logger.Log.Println(
			"HOSTS DISCOVERED:",
			len(network.Hosts),
		)

		network.Hosts = EnrichHosts(
			network.Hosts,
			iface,
		)

		logger.Log.Println(
			"HOSTS ENRICHED:",
			len(network.Hosts),
		)

		networks = append(
			networks,
			network,
		)
	}

	return networks
}

// Topology builds a graph representation
// of the discovered network.
//
// It performs:
//
// Scan()
//     ↓
// BuildTopology()
//     ↓
// Ping every node
//     ↓
// Update link latency and status

func (s *Scanner) Topology() models.Topology {

	networks := s.Scan()

	topology := BuildTopology(
		networks,
	)

	pingResults := make(
		map[string]models.Host,
	)

	for i := range topology.Nodes {

		node := &topology.Nodes[i]

		result := PingHost(
			node.IP,
			2*time.Second,
		)

		pingResults[node.IP] = result

		node.Online = result.Online
		node.RTT = result.RTT
	}

	for i := range topology.Links {

		result :=
			pingResults[topology.Links[i].To]

		if result.Online {

			latency :=
				result.RTT.Seconds() * 1000

			topology.Links[i].Latency = latency

			switch {

			case latency < 10:
				topology.Links[i].Status = "good"

			case latency < 50:
				topology.Links[i].Status = "warning"

			default:
				topology.Links[i].Status = "critical"
			}

		} else {

			topology.Links[i].Latency = 0

			topology.Links[i].Status = "timeout"
		}
	}

	return topology
}
