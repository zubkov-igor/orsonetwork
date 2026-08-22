package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func EnrichMDNS(
  hosts []models.Host,
) []models.Host {

  logger.Log.Println(
    "MDNS ENRICHMENT START",
  )

  mdnsRecords := DiscoverMDNS()

  logger.Log.Println(
    "MDNS FOUND:",
    len(mdnsRecords),
  )



  for i := range hosts {

    for _, mdns := range mdnsRecords {

    	 logger.Log.Println(
    "MDNS COMPARE:",
    "host=", hosts[i].IP,
    "mdns=", mdns.IP,
)

      if mdns.IP != hosts[i].IP {
        continue
      }

      hosts[i].MDNS = append(
        hosts[i].MDNS,
        mdns,
      )

      logger.Log.Println(
        "MDNS MATCH:",
        hosts[i].IP,
        mdns.Name,
        mdns.Service,
        mdns.Host,
        mdns.Port,
      )
    }
  }

  logger.Log.Println(
    "MDNS ENRICHMENT FINISHED",
  )

  return hosts
}
