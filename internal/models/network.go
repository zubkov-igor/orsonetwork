package models

type Network struct {
	CIDR      string `json:"cidr"`
	Interface string `json:"interface"`
	Gateway   string `json:"gateway"`
	Hosts     []Host `json:"hosts"`
}
