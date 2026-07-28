package models

type DiscoveryType string

const (
    DiscoveryARP        DiscoveryType = "arp"
    DiscoveryReverseDNS DiscoveryType = "reverse_dns"
    DiscoveryNetBIOS    DiscoveryType = "netbios"
    DiscoveryMDNS DiscoveryType = "mdns"
)