package models

import "time"

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`

	IP       string `json:"ip"`
	MAC      string `json:"mac"`
	Hostname string `json:"hostname"`
	Vendor   string `json:"vendor"`

	Sources []string `json:"sources"`

	Online bool          `json:"online"`
	RTT    time.Duration `json:"rtt"`
}
