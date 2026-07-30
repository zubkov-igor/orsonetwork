package models

import "time"

// Host is the main entity produced by discovery pipeline.
// It contains raw discovery data, enrichment results,
// and classification information.

type Host struct {
	IP string `json:"ip"`

	MAC string `json:"mac"`

	Hostname string `json:"hostname"`

	Vendor string `json:"vendor"`

	Ports []Port `json:"ports"`

	MDNS []MDNSService `json:"mdns"`

	UDPServices []UDPService `json:"udp"`

	Type DeviceType `json:"deviceType"`

	Sources []DiscoverySource `json:"sources"`

	Online bool

	RTT time.Duration
}
