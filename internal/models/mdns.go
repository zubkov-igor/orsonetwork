package models

// MDNSService represents a discovered multicast DNS service.
// Despite the name, it also stores host-level discovery data.

type MDNSService struct {
	Name    string   `json:"name"`
	Service string   `json:"service"`
	Host    string   `json:"host"`
	IP      string   `json:"ip"`
	Port    int      `json:"port"`
	TXT     []string `json:"txt"`
}
