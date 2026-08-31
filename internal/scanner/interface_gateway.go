package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func GatewayForInterface(
	iface models.Interface,
	gateways []models.Gateway,
) *models.Gateway {

	logger.Debug(
		"GATEWAY SEARCH FOR INTERFACE:",
		iface.Name,
	)

	for _, gw := range gateways {

		logger.Debug(
			"CHECK GATEWAY:",
			gw.IP,
			gw.Interface,
		)

		if gw.Interface == iface.Name {

			logger.Debug(
				"GATEWAY MATCH:",
				gw.IP,
			)

			return &gw
		}
	}

	for i := range gateways {

		logger.Debug(
			"CHECK GATEWAY:",
			gateways[i].IP,
			gateways[i].Interface,
		)

		if gateways[i].Interface == iface.Name {

			logger.Info(
				"GATEWAY MATCH:",
				gateways[i].IP,
			)

			return &gateways[i]
		}
	}

	logger.Warn(
		"GATEWAY NOT FOUND",
	)

	return nil
}
