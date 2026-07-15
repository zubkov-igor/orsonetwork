package scanner

import "OrsoNetwork/internal/models"

func BuildNetwork(
    iface models.Interface,
    gateway *models.Gateway,
) models.Network {

    network := GetNetwork(iface)

    if network == nil {
        return models.Network{}
    }

    return models.Network{
        CIDR:      network.String(),
        Interface: iface.Name,
        Gateway:   gateway.IP,
    }
}