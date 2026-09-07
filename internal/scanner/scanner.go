package scanner

import (
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

type Scanner struct {
	Config models.ScannerConfig
}

func New() *Scanner {

	config, err := LoadConfig()

	if err != nil {

		logger.Error(
			"CONFIG LOAD FAILED:",
			err,
		)

		config = DefaultScannerConfig()
	}

	logger.Debug(
		"SCANNER CONFIG LOADED:",
		config,
	)
	return &Scanner{
		Config: config,
	}
}

func (s *Scanner) UpdateConfig(
	config models.ScannerConfig,
) error {

	s.Config = config

	return SaveConfig(
		config,
	)
}

func (s *Scanner) GetConfig() models.ScannerConfig {
	return s.Config
}

func DefaultScannerConfig() models.ScannerConfig {

	return models.ScannerConfig{

		EnableICMP: true,

		EnableARP: true,

		EnableReverseDNS: true,

		EnableNetBIOS: true,

		EnableMDNS: true,

		EnableSSDP: true,

		EnableSNMP: false,

		EnableTCP: true,

		EnableUDP: true,

		Workers: 20,
	}
}

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

	logger.Info(
		"SCANNER INTERFACES:",
		len(interfaces),
	)

	gateways := GetGateways()

	logger.Info(
		"SCANNER GATEWAYS:",
		len(gateways),
	)

	logger.Debug(
		"SCANNER LOOP START",
	)

	for _, iface := range interfaces {

		logger.Debug(
			"PROCESS INTERFACE:",
			iface.Name,
		)

		gw := GatewayForInterface(
			iface,
			gateways,
		)

		if gw == nil {

			logger.Warn(
				"NO GATEWAY FOR INTERFACE:",
				iface.Name,
			)

			continue
		}

		logger.Info(
			"GATEWAY FOUND:",
			iface.Name,
			gw.IP,
		)

		network := BuildNetwork(
			iface,
			gw,
		)

		logger.Info(
			"NETWORK BUILT:",
			network.CIDR,
		)

		subnets := SubnetsFromCIDR(
			network.CIDR,
			24,
		)

		logger.Info(
			"SUBNETS GENERATED:",
			len(subnets),
		)

		for _, subnet := range subnets {

			logger.Debug(
				"SUBNET:",
				subnet,
			)
		}

		// TODO: временно отключено до тестирования больших сетей.
		//
		// if len(subnets) > 1 {
		//
		//     logger.Info(
		//         "LARGE NETWORK DETECTED:",
		//         network.CIDR,
		//     )
		//
		//     continue
		// }

		ips := HostsFromCIDR(
			network.CIDR,
		)
		logger.Info(
			"HOST IPS GENERATED:",
			len(ips),
		)

		network.Hosts = DiscoverHostsFull(
			ips,
			s.Config.Workers,
			s.Config,
		)

		logger.Info(
			"HOSTS DISCOVERY FULL FINISHED:",
			len(network.Hosts),
		)

network.Hosts = EnrichHosts(
    network.Hosts,
    iface,
    s.Config,
)

for i := range network.Hosts {

    host := &network.Hosts[i]

    evidence := BuildFingerprint(*host)

    host.Fingerprint = evidence

    host.OS = IdentifyOS(*host)

    logger.Info(
        "OS IDENTIFIED:",
        host.IP,
        "OS:",
        host.OS,
    )
}

logger.Info(
    "HOSTS ENRICHMENT FINISHED:",
    len(network.Hosts),
)

var ssdpResponses []SSDPResponse

if s.Config.EnableSSDP {

    ssdpResponses = ProbeSSDP(
        iface,
    )

    logger.Info(
        "SSDP RESPONSES:",
        len(ssdpResponses),
    )
}

if s.Config.EnableMDNS {

    mdnsServices := DiscoverMDNS()

    logger.Info(
        "MDNS DISCOVERED:",
        len(mdnsServices),
    )
}

networks = append(
    networks,
    network,
)
}

	return networks
}


func (s *Scanner) Topology() models.ScanResult {

	start := time.Now()

	networks := s.Scan()

	topology := BuildTopology(
		networks,
	)

	logger.Info(
		"TOPOLOGY NODES:",
		len(topology.Nodes),
	)

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

		if !result.Online {

			topology.Links[i].Latency = 0
			topology.Links[i].Status = "timeout"

		} else if result.RTT == 0 {

			topology.Links[i].Latency = 0
			topology.Links[i].Status = "unknown"

		} else {

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
		}
	}

	for _, node := range topology.Nodes {

		logger.Debug(
			"FINAL NODE:",
			node.IP,
			"ONLINE:",
			node.Online,
			"RTT:",
			node.RTT,
		)
	}

	for _, link := range topology.Links {

		logger.Debug(
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
