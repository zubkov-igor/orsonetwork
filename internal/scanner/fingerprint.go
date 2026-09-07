package scanner

import (
	"fmt"
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
	"strings"
)

func BuildFingerprint(host models.Host) []models.FingerprintEvidence {

	var evidence []models.FingerprintEvidence

	hostname := strings.ToLower(host.Hostname)

	if hostname != "" {
		evidence = append(
			evidence,
			models.FingerprintEvidence{
				Source: models.DiscoveryReverseDNS,
				Value:  "hostname:" + hostname,
			},
		)
	}

	vendor := strings.ToLower(host.Vendor)

	if vendor != "" {
		evidence = append(
			evidence,
			models.FingerprintEvidence{
				Source: models.DiscoveryOUI,
				Value:  "vendor:" + vendor,
			},
		)
	}

	model, ok := FindDeviceModelFromHostname(host.Hostname)

	if ok {
		evidence = append(
			evidence,
			models.FingerprintEvidence{
				Source: models.DiscoveryReverseDNS,
				Value:  "model:" + model.Brand + " " + model.Name,
			},
		)
	}

	// =========================
	// TCP
	// =========================

	logger.Debug(
		"TCP FINGERPRINT INPUT:",
		host.IP,
		host.Ports,
	)

	for _, port := range host.Ports {

    if !port.Open {
        continue
    }

    value := fmt.Sprintf(
        "tcp-port:%d",
        port.Number,
    )

    if port.Service != "" {
        value += ":" + strings.ToLower(port.Service)
    }

    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  value,
        },
    )

    if port.Banner != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryTCP,
                Value:  "tcp-banner:" + port.Banner,
            },
        )
    }

		switch port.Number {

		case 22:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "ssh",
			})
		case 135:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "msrpc",
			})

		case 139:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "netbios",
			})

		case 445:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "smb",
			})

		case 3389:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "rdp",
			})
		case 7680:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "port:7680",
			})

		case 554:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "rtsp",
			})

		case 631:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "ipp",
			})
		case 9100:
			evidence = append(evidence, models.FingerprintEvidence{
				Source: models.DiscoveryTCP,
				Value:  "raw-printing",
			})
		}
	}

	// =========================
	// UDP
	// =========================

for _, service := range host.UDPServices {

    if service.Port != 0 {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryUDP,
                Value:  fmt.Sprintf("udp-port:%d", service.Port),
            },
        )
    }

    if service.Protocol != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryUDP,
                Value:  "udp-protocol:" + strings.ToLower(service.Protocol),
            },
        )
    }

    if service.Service != "" {
        source := models.DiscoveryUDP

        if service.Port == 161 {
            source = models.DiscoverySNMP
        }

        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: source,
                Value:  "service:" + strings.ToLower(service.Service),
            },
        )
    }

    if service.Info != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryUDP,
                Value:  "udp-info:" + service.Info,
            },
        )
    }
}

	// =========================
	// mDNS
	// =========================

for _, service := range host.MDNS {

    if service.Name != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value:  "name:" + service.Name,
            },
        )
    }

    if service.Service != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value:  "service:" + service.Service,
            },
        )
    }

    if service.Host != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value:  "host:" + service.Host,
            },
        )
    }

    if service.IP != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value:  "ip:" + service.IP,
            },
        )
    }

    if service.Port != 0 {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value: fmt.Sprintf(
                    "port:%d",
                    service.Port,
                ),
            },
        )
    }

    for _, txt := range service.TXT {
        if txt == "" {
            continue
        }

        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoveryMDNS,
                Value: "txt:" + txt,
            },
        )
    }
}
	for _, httpInfo := range host.HTTP {

		if httpInfo.Server != "" {
			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryTCP,
					Value:  "http-server:" + httpInfo.Server,
				},
			)
		}

		if httpInfo.Title != "" {
			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryTCP,
					Value:  "http-title:" + httpInfo.Title,
				},
			)
		}

if httpInfo.StatusCode != 0 {
    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  fmt.Sprintf("http-status:%d", httpInfo.StatusCode),
        },
    )
}

if httpInfo.ContentType != "" {
    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  "http-content-type:" + httpInfo.ContentType,
        },
    )
}

for _, keyword := range httpInfo.Keywords {
    if keyword == "" {
        continue
    }

    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  "http-keyword:" + keyword,
        },
    )
}

for _, script := range httpInfo.Scripts {
    if script == "" {
        continue
    }

    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  "http-script:" + script,
        },
    )
}

for _, fingerprint := range httpInfo.Fingerprint {
    if fingerprint == "" {
        continue
    }

    evidence = append(
        evidence,
        models.FingerprintEvidence{
            Source: models.DiscoveryTCP,
            Value:  "http-fingerprint:" + fingerprint,
        },
    )
}
}

	for _, source := range host.Sources {
		if source.Type == models.DiscoveryNetBIOS &&
			source.Value != "" {

			evidence = append(
				evidence,
				models.FingerprintEvidence{
					Source: models.DiscoveryNetBIOS,
					Value:  "name:" + source.Value,
				},
			)
		}
	}


	for _, info := range host.SNMP {

    if info.Version != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoverySNMP,
                Value:  "snmp-version:" + info.Version,
            },
        )
    }

    if info.SysDescr != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoverySNMP,
                Value:  "snmp-sysdescr:" + info.SysDescr,
            },
        )
    }

    if info.SysName != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoverySNMP,
                Value:  "snmp-sysname:" + info.SysName,
            },
        )
    }

    if info.SysLocation != "" {
        evidence = append(
            evidence,
            models.FingerprintEvidence{
                Source: models.DiscoverySNMP,
                Value:  "snmp-syslocation:" + info.SysLocation,
            },
        )
    }
}

	return evidence
}
