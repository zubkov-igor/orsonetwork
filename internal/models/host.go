package models

import "time"

type Host struct {
	IP       string `json:"ip"`
	MAC      string `json:"mac"`
	Hostname string `json:"hostname"`
	Vendor   string `json:"vendor"`

	Sources []DiscoverySource `json:"sources"`

	Online bool          `json:"online"`
	RTT    time.Duration `json:"rtt"`
}
