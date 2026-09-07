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
      OS       string `json:"os"`

      Ports       []Port             `json:"ports"`
      HTTP        []HTTPInfo         `json:"http"`
      MDNS        []MDNSService      `json:"mdns"`
      UDPServices []UDPService       `json:"udpServices"`
      SNMP        []SNMPInfo         `json:"snmp"`
      Sources     []DiscoverySource  `json:"sources"`
      Fingerprint []FingerprintEvidence `json:"fingerprint"`

      Online bool          `json:"online"`
      RTT    time.Duration `json:"rtt"`
}