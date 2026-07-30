package models

// Link represents a logical connection
// between two network nodes.
// Link information may be discovered
// from ARP, SNMP, LLDP or topology analysis.

type Link struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Type    string  `json:"type"`
	Latency float64 `json:"latency"`
	Status  string  `json:"status"`
}
