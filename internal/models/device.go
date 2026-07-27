package models

type Device struct {
	IP       string
	MAC      string
	Hostname string
	Vendor string
	Ports    []Port
}