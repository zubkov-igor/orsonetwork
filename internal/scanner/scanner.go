package scanner

import (
    "time"

    "OrsoNetwork/internal/models"
)

type Scanner struct {
}

func New() *Scanner {
    return &Scanner{}
}

func (s *Scanner) Scan() []models.Network {

    var networks []models.Network

    interfaces := GetInterfaces()
    gateways := GetGateways()

    for _, iface := range interfaces {

        gw := GatewayForInterface(
            iface.Name,
            gateways,
        )

        if gw == nil {
            continue
        }

        network := BuildNetwork(
            iface,
            gw,
        )

network.Hosts = DiscoverHosts(
    network.CIDR,
)

network.Hosts = EnrichHosts(
    network.Hosts,
)

for _, h := range network.Hosts {

    println(
        "SCAN HOST:",
        h.IP,
        h.MAC,
        h.Vendor,
    )
}

        networks = append(
            networks,
            network,
        )
    }

    return networks
}

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

            topology.Links[i].Latency =
                result.RTT.Seconds() * 1000

        } else {

            topology.Links[i].Latency = 0
        }
    }

    return topology
}