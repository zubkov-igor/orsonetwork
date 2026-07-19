package scanner

import (
    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func GatewayForInterface(
    iface models.Interface,
    gateways []models.Gateway,
) *models.Gateway {

    logger.Log.Println(
        "GATEWAY SEARCH FOR INTERFACE:",
        iface.Name,
    )

    for _, gw := range gateways {

        logger.Log.Println(
            "CHECK GATEWAY:",
            gw.IP,
            gw.Interface,
        )

        if gw.Interface == iface.Name {

            logger.Log.Println(
                "GATEWAY MATCH:",
                gw.IP,
            )

            return &gw
        }
    }

    logger.Log.Println(
        "GATEWAY NOT FOUND",
    )

    return nil
}