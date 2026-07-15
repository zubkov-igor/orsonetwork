package models

type Network struct {
	CIDR      string
	Interface string
	Gateway   string
	Hosts     []Host
}