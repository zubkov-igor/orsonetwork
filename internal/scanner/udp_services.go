package scanner

// UDPServices contains known UDP ports and their service names.
// Used for UDP discovery and later fingerprinting.
var UDPServices = map[int]string{

	137:  "netbios",
	161:  "snmp",
	1900: "ssdp",
	5353: "mdns",
}
