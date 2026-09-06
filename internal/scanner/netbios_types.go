package scanner

type NetBIOSName struct {
	Name   string
	Suffix byte
	Flags  uint16
}

// NetBIOSResult contains information
// extracted from NetBIOS discovery.

type NetBIOSResult struct {
	Name      string
	Workgroup string
	MAC       string
	Names     []NetBIOSName
}
