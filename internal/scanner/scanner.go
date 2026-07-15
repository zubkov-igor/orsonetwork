package scanner

import "OrsoNetwork/internal/models"

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

        networks = append(
            networks,
            network,
        )
    }

    return networks
}

func (s *Scanner) Topology() models.Topology {

    networks := s.Scan()

    return BuildTopology(
        networks,
    )
}