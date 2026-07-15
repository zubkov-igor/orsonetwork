package scanner

import (
	"net"

	"OrsoNetwork/internal/models"
)

func GetNetwork(iface models.Interface) *net.IPNet {

	ip := net.ParseIP(iface.IP)

	if ip == nil {
		return nil
	}

	mask := net.CIDRMask(24, 32)

	network := &net.IPNet{
		IP:   ip.Mask(mask),
		Mask: mask,
	}

	return network
}