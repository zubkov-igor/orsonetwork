package scanner

import "OrsoNetwork/internal/models"

// BuildNetwork creates a network model
// for a single interface.
//
// It combines:
//
// - interface information
// - network CIDR
// - gateway
//
// Hosts are discovered later.

func BuildNetwork(
	iface models.Interface,
	gateway *models.Gateway,
) models.Network {

	// Calculate network address
	// from interface IP and mask.

	network := GetNetwork(iface)

	// Interface has no valid network.

	if network == nil {
		return models.Network{}
	}

	return models.Network{
		CIDR:      network.String(),
		Interface: iface.Name,
		Gateway:   gateway.IP,
		Hosts:     []models.Host{},
	}
}
