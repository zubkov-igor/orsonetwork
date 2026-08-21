package models

import "time"

// Host is the main entity produced by discovery pipeline.
// It contains raw discovery data, enrichment results,
// and classification information.

type Host struct {
	IP       string
	MAC      string
	Hostname string
	Vendor   string

	Ports []Port

	HTTP []HTTPInfo

	MDNS []MDNSService

	UDPServices []UDPService

	SNMP []SNMPInfo

	Type DeviceType

	Confidence int

	Sources []DiscoverySource

	Online bool

	RTT time.Duration
}
