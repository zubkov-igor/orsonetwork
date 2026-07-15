package scanner

import "OrsoNetwork/internal/models"

func GatewayForInterface(
	iface string,
	gateways []models.Gateway,
) *models.Gateway {

	for _, gw := range gateways {

		if gw.Interface == iface {
			return &gw
		}

	}

	return nil
}