package models

// DiscoveryType describes
// how information about a host was discovered.
type DiscoveryType string

const (
    DiscoveryARP        DiscoveryType = "arp"
    DiscoveryOUI        DiscoveryType = "oui"

    DiscoveryReverseDNS DiscoveryType = "reverse_dns"
    DiscoveryNetBIOS    DiscoveryType = "netbios"
    DiscoveryMDNS       DiscoveryType = "mdns"

    DiscoverySNMP       DiscoveryType = "snmp"
    DiscoverySSDP       DiscoveryType = "ssdp"
)