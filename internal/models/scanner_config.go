package models

type ScannerConfig struct {
    EnableICMP       bool `json:"enable_icmp"`
    EnableARP        bool `json:"enable_arp"`
    EnableReverseDNS bool `json:"enable_reverse_dns"`
    EnableNetBIOS    bool `json:"enable_netbios"`
    EnableMDNS       bool `json:"enable_mdns"`
    EnableSSDP       bool `json:"enable_ssdp"`
    EnableSNMP       bool `json:"enable_snmp"`
    EnableTCP        bool `json:"enable_tcp"`
    EnableUDP        bool `json:"enable_udp"`

    Workers int `json:"workers"`
}