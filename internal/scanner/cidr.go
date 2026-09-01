package scanner

import (
	"net"
)

func HostsFromCIDR(cidr string) []string {

	var hosts []string

	ip, network, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil
	}

	for ip := ip.Mask(network.Mask); network.Contains(ip); incIP(ip) {

		if ip.Equal(network.IP) {
			continue
		}

		hosts = append(
			hosts,
			ip.String(),
		)
	}

	if len(hosts) > 0 {

		hosts = hosts[:len(hosts)-1]
	}

	return hosts
}

func incIP(ip net.IP) {

	for j := len(ip) - 1; j >= 0; j-- {

		ip[j]++

		if ip[j] != 0 {
			break
		}
	}
}


func SubnetsFromCIDR(cidr string, prefix int) []string {

	ip, network, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil
	}

	ones, bits := network.Mask.Size()

	if prefix < ones || prefix > bits {
		return nil
	}

	if prefix == ones {
		return []string{
			network.String(),
		}
	}

	var subnets []string

	subnetMask := net.CIDRMask(
		prefix,
		bits,
	)

	subnetSize := uint32(1) << uint(bits-prefix)

	baseIP := ip.Mask(network.Mask)

	for offset := uint32(0); ; offset += subnetSize {

		subnetIP := make(net.IP, len(baseIP))
		copy(subnetIP, baseIP)

		addIP(subnetIP, offset)

		subnet := net.IPNet{
			IP:   subnetIP,
			Mask: subnetMask,
		}

		if !network.Contains(subnetIP) {
			break
		}

		subnets = append(
			subnets,
			subnet.String(),
		)
	}

	return subnets
}

func addIP(ip net.IP, value uint32) {

	for i := len(ip) - 1; i >= 0 && value > 0; i-- {

		sum := uint32(ip[i]) + (value & 0xff)

		ip[i] = byte(sum)

		value = (value >> 8) + (sum >> 8)
	}
}