package scanner

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

// EnrichHosts collects additional information
// for every discovered host.
//
// Discovery pipeline:
//
// mDNS
// ARP
// UDP
// Reverse DNS
// NetBIOS
// Port Scan
// Device Identification

func EnrichHosts(
	hosts []models.Host,
) []models.Host {

	// Discover mDNS services
	mdnsRecords := DiscoverMDNS()

	logger.Log.Println(
		"MDNS FOUND:",
		len(mdnsRecords),
	)

	// Resolve MAC addresses using ARP
	arpHosts := ARPDiscovery(
		hosts,
	)

	// Build fast lookup table:
	// IP -> ARP result

	arpMap := make(
		map[string]models.Host,
	)

	for _, h := range arpHosts {
		arpMap[h.IP] = h
	}

	for i := range hosts {

		// ==========================
		// mDNS enrichment
		// ==========================

		for _, mdns := range mdnsRecords {

			if mdns.IP == hosts[i].IP {

				hosts[i].MDNS = append(
					hosts[i].MDNS,
					mdns,
				)

				hosts[i].Sources = append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryMDNS,
						Value: mdns.Name,
					},
				)
			}
		}

		// ==========================
		// ARP enrichment
		// MAC + Vendor
		// ==========================

		if arpHost, ok := arpMap[hosts[i].IP]; ok {

			if arpHost.MAC != "" {

				hosts[i].MAC = arpHost.MAC

				hosts[i].Sources = append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryARP,
						Value: arpHost.MAC,
					},
				)

				hosts[i].Vendor = LookupVendor(
					arpHost.MAC,
				)

				if hosts[i].Vendor != "" {

					hosts[i].Sources = append(
						hosts[i].Sources,
						models.DiscoverySource{
							Type:  models.DiscoveryOUI,
							Value: hosts[i].Vendor,
						},
					)
				}
			}
		}

		// ==========================
		// UDP Discovery
		// ==========================

		hosts[i].UDPServices = DiscoverUDP(
			hosts[i].IP,
		)

		for _, u := range hosts[i].UDPServices {

			logger.Log.Println(
				"UDP SERVICE:",
				hosts[i].IP,
				u.Port,
				u.Service,
			)
		}

		// ==========================
		// Reverse DNS
		// ==========================

		hostname := LookupReverseDNS(
			hosts[i].IP,
		)

		if hostname != "" {

			hosts[i].Hostname = hostname

			hosts[i].Sources = append(
				hosts[i].Sources,
				models.DiscoverySource{
					Type:  models.DiscoveryReverseDNS,
					Value: hostname,
				},
			)
		}

		// ==========================
		// NetBIOS
		// ==========================

		netbios, err := LookupNetBIOS(
			hosts[i].IP,
		)

		if err == nil {

			logger.Log.Println(
				"NETBIOS FOUND:",
				hosts[i].IP,
				netbios.Name,
				netbios.MAC,
			)

			if netbios.Name != "" {

				if hosts[i].Hostname == "" {

					hosts[i].Hostname =
						netbios.Name
				}

				hosts[i].Sources = append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryNetBIOS,
						Value: netbios.Name,
					},
				)
			}

			// NetBIOS MAC fallback

			if hosts[i].MAC == "" &&
				netbios.MAC != "" {

				hosts[i].MAC =
					netbios.MAC

				hosts[i].Vendor =
					LookupVendor(
						netbios.MAC,
					)

				hosts[i].Sources = append(
					hosts[i].Sources,
					models.DiscoverySource{
						Type:  models.DiscoveryNetBIOS,
						Value: netbios.MAC,
					},
				)
			}
		}

		// ==========================
		// TCP Port scan
		// ==========================

		hosts[i].Ports = ScanPorts(
			hosts[i].IP,
		)

		for _, p := range hosts[i].Ports {

			logger.Log.Println(
				"OPEN PORT:",
				hosts[i].IP,
				p.Number,
				p.Protocol,
				p.Service,
			)
		}

		// ==========================
		// Device identification
		// ==========================

		hosts[i].Type =
			IdentifyDevice(
				hosts[i],
			)

		logger.Log.Println(
			"DEVICE DETECTED:",
			hosts[i].IP,
			hosts[i].Hostname,
			hosts[i].Vendor,
			hosts[i].Type,
		)
	}

	return hosts
}
