package scanner

import (
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func discoverHostNew(
    ip string,
    config models.ScannerConfig,
) models.Host {


    host := models.Host{
        IP: ip,
    }
  start := time.Now()

  // Ping

  stepStart := time.Now()

  result := PingHost(
    ip,
    1*time.Millisecond,
  )

  pingDuration := time.Since(stepStart)

  host.Online = result.Online
  host.RTT = result.RTT

  if pingDuration > 100*time.Millisecond {

    logger.Log.Println(
      "DISCOVERY SLOW:",
      ip,
      "PING:",
      pingDuration,
    )
  }

  if result.Online {

    host.Sources = append(
      host.Sources,
      models.DiscoverySource{
        Type:  models.DiscoveryICMP,
        Value: result.RTT.String(),
      },
    )
  }

 // ARP

if config.EnableARP {

    stepStart = time.Now()

    mac := ARPResolve(ip)

    arpDuration := time.Since(stepStart)

    if arpDuration > 100*time.Millisecond {

        logger.Log.Println(
            "DISCOVERY SLOW:",
            ip,
            "ARP:",
            arpDuration,
        )
    }

    if mac != "" {

        host.MAC = mac

        host.Vendor = LookupVendor(mac)

        host.Sources = append(
            host.Sources,
            models.DiscoverySource{
                Type:  models.DiscoveryARP,
                Value: mac,
            },
        )

        if host.Vendor != "" && host.Vendor != "Unknown" {

            host.Sources = append(
                host.Sources,
                models.DiscoverySource{
                    Type:  models.DiscoveryOUI,
                    Value: host.Vendor,
                },
            )
        }
    }
}

// NetBIOS

if result.Online && config.EnableNetBIOS {

    stepStart = time.Now()

    data, err := ProbeNetBIOS(ip)

    netbiosDuration := time.Since(stepStart)

    if netbiosDuration > 100*time.Millisecond {

        logger.Log.Println(
            "DISCOVERY SLOW:",
            ip,
            "NETBIOS:",
            netbiosDuration,
        )
    }

    if err == nil {

        netbios := ParseNetBIOSResponse(data)

        if netbios.Name != "" {

            host.Hostname = netbios.Name

            host.Sources = append(
                host.Sources,
                models.DiscoverySource{
                    Type:  models.DiscoveryNetBIOS,
                    Value: netbios.Name,
                },
            )
        }

        if netbios.MAC != "" && host.MAC == "" {
            host.MAC = netbios.MAC
        }
    }
}


// Reverse DNS

if config.EnableReverseDNS {

    stepStart = time.Now()

    hostname := LookupReverseDNS(ip)

    dnsDuration := time.Since(stepStart)

    if dnsDuration > 100*time.Millisecond {

        logger.Log.Println(
            "DISCOVERY SLOW:",
            ip,
            "REVERSE DNS:",
            dnsDuration,
        )
    }

    if hostname != "" && host.Hostname == "" {

        host.Hostname = hostname

        host.Sources = append(
            host.Sources,
            models.DiscoverySource{
                Type:  models.DiscoveryReverseDNS,
                Value: hostname,
            },
        )
    }
}
  totalDuration := time.Since(start)

  if totalDuration > 1*time.Second {

    logger.Log.Println(
      "DISCOVERY SLOW TOTAL:",
      ip,
      totalDuration,
    )
  }

identification := IdentifyDevice(
    host,
)

host.Type = identification.Type

host.Confidence =
    CalculateConfidence(
        host,
    )

logger.Log.Println(
    "IDENTIFICATION:",
    host.IP,
    "TYPE:",
    host.Type,
    "CONFIDENCE:",
    host.Confidence,
    "MAC:",
    host.MAC,
    "VENDOR:",
    host.Vendor,
    "HOSTNAME:",
    host.Hostname,
)

return host

}