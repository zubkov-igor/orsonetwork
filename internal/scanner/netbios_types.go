package scanner

// NetBIOSResult contains information
// extracted from NetBIOS discovery.

type NetBIOSResult struct {
	Name      string
	Workgroup string
	MAC       string
}
