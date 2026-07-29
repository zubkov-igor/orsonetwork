package models

import "time"

// Node represents a network graph vertex.
// It is created from discovered Host data
// and used for topology visualization.

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`

	IP       string `json:"ip"`
	MAC      string `json:"mac"`
	Hostname string `json:"hostname"`
	Vendor   string `json:"vendor"`

	Sources []DiscoverySource `json:"sources"`

	Online bool          `json:"online"`
	RTT    time.Duration `json:"rtt"`
}
