package scanner

import (
	
	"OrsoNetwork/internal/models"
	"OrsoNetwork/internal/logger"
)

func enrichHostMDNS(
    host models.Host,
    mdnsRecords []models.MDNSService,
) models.Host {

    for _, mdns := range mdnsRecords {

        if mdns.IP != host.IP {
            continue
        }

        host.MDNS = append(
            host.MDNS,
            mdns,
        )

        logger.Log.Println(
            "MDNS MATCH:",
            host.IP,
            mdns.Name,
            mdns.Service,
            mdns.Host,
            mdns.Port,
        )
    }

    return host
}