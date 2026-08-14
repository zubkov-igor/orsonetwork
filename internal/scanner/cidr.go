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
