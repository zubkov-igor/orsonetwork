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

		ips := HostsFromCIDR(
	network.CIDR,
)

logger.Log.Println(
	"HOST IPS GENERATED:",
	len(ips),
)

network.Hosts = DiscoverHostsFull(
    ips,
    20,
)

logger.Log.Println(
    "HOSTS DISCOVERY FULL FINISHED:",
    len(network.Hosts),
)

network.Hosts = EnrichHosts(
    network.Hosts,
    iface,
)

logger.Log.Println(
    "HOSTS ENRICHMENT FINISHED:",
    len(network.Hosts),
)

		mdnsIPs := ProbeMDNS(
			iface,
		)

		ssdpResponses := ProbeSSDP(iface)

logger.Log.Println(
    "SSDP RESPONSES:",
    len(ssdpResponses),
)

for _, response := range ssdpResponses {

    logger.Log.Println(
        "SSDP DEVICE:",
        response.IP,
        "LOCATION:",
        response.Location,
        "SERVER:",
        response.Server,
        "ST:",
        response.ST,
        "USN:",
        response.USN,
    )
}

		logger.Log.Println(
			"MDNS IPS:",
			mdnsIPs,
		)

		// Merge mDNS discovered hosts
		// with hosts discovered by ICMP.

		knownHosts := make(
			map[string]bool,
		)

		for _, host := range network.Hosts {

			knownHosts[host.IP] = true
		}

		for _, ip := range mdnsIPs {

			// Never add our own host.

			if ip == iface.IP {
				continue
			}

			// Host already discovered by ICMP.

			if knownHosts[ip] {
				continue
			}

			network.Hosts = append(
				network.Hosts,
				models.Host{
					IP: ip,
				},
			)

			knownHosts[ip] = true
		}

		logger.Log.Println(
			"HOSTS AFTER MDNS MERGE:",
			len(network.Hosts),
		)

		for _, host := range network.Hosts {

			logger.Log.Println(
				"HOST:",
				host.IP,
			)
		}

		logger.Log.Println(
			"HOSTS DISCOVERED:",
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

func (s *Scanner) Topology() models.ScanResult {

	start := time.Now()

	networks := s.Scan()

	topology := BuildTopology(
		networks,
	)

	logger.Log.Println(
	"TOPOLOGY NODES:",
	len(topology.Nodes),
)

for _, node := range topology.Nodes {

	logger.Log.Println(
		"TOPOLOGY NODE:",
		node.IP,
		"TYPE:",
		node.Type,
		"MAC:",
		node.MAC,
	)
}

	nodeIPs := make(map[string]string)

	for _, node := range topology.Nodes {
		nodeIPs[node.ID] = node.IP
	}

	for i := range topology.Links {

		ip := nodeIPs[topology.Links[i].To]

		var result models.Host

		for _, node := range topology.Nodes {

			if node.IP == ip {

				result.Online = node.Online
				result.RTT = node.RTT

				break
			}
		}

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

	for _, node := range topology.Nodes {

		logger.Log.Println(
			"FINAL NODE:",
			node.IP,
			"ONLINE:",
			node.Online,
			"RTT:",
			node.RTT,
		)
	}

	for _, link := range topology.Links {

		logger.Log.Println(
			"FINAL LINK:",
			link.From,
			"->",
			link.To,
			"LATENCY:",
			link.Latency,
			"STATUS:",
			link.Status,
		)
	}

	return models.ScanResult{
		Topology: topology,
		Duration: time.Since(start).Milliseconds(),
		LastScan: time.Now().Unix(),
	}
}

