package scanner

import (
	"net"
	"net/netip"
	"time"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"

	"github.com/mdlayher/arp"
)

func ARPDiscovery(
	hosts []models.Host,
) []models.Host {

	logger.Debug(
		"ARP DISCOVERY START",
	)

	var arpHosts []models.Host

	var interfaceIP string

	interfaces := GetInterfaces()

	for _, i := range interfaces {

		logger.Debug(
			"AVAILABLE INTERFACE:",
			i.Name,
		)
	}

	if len(interfaces) == 0 {
		return arpHosts
	}

	var iface *net.Interface

	for _, i := range interfaces {

		if i.Name == "" {
			continue
		}

		found, err := net.InterfaceByName(
			i.Name,
		)

		if err != nil {
			continue
		}

		if IsVirtualInterface(found.Name) {

			logger.Debug(
				"SKIP INTERFACE:",
				found.Name,
			)

			continue
		}

		iface = found
		interfaceIP = i.IP

		logger.Debug(
			"ARP INTERFACE:",
			iface.Name,
			iface.HardwareAddr.String(),
		)

		break
	}

	if iface == nil {
		logger.Warn(
			"NO ARP INTERFACE",
		)

		return arpHosts
	}

	client, err := arp.Dial(
		iface,
	)

	if err != nil {

		logger.Warn(
			"ARP DIAL ERROR:",
			err.Error(),
		)

		return arpHosts
	}

	defer client.Close()

	for _, host := range hosts {

		if host.IP == interfaceIP {
			logger.Debug(
				"SKIP OWN HOST:",
				host.IP,
			)

			continue
		}

		addr, err := netip.ParseAddr(
			host.IP,
		)

		if err != nil {
			continue
		}

		logger.Debug(
			"ARP REQUEST:",
			host.IP,
		)

		err = client.SetReadDeadline(
			time.Now().Add(
				1 * time.Second,
			),
		)

		if err != nil {
			continue
		}

		mac, err := client.Resolve(
			addr,
		)

		if err != nil {

			logger.Warn(
				"ARP RESOLVE ERROR:",
				host.IP,
				err.Error(),
			)

			arpHosts = append(
				arpHosts,
				host,
			)

			continue
		}

		arpHost := host

		arpHost.MAC = mac.String()

		logger.Info(
			"ARP HOST:",
			arpHost.IP,
			arpHost.MAC,
		)

		arpHosts = append(
			arpHosts,
			arpHost,
		)
	}

	return arpHosts
}

func ARPResolve(ip string) string {

	interfaces := GetInterfaces()

	var iface *net.Interface

	for _, i := range interfaces {

		if i.Name == "" {
			continue
		}

		found, err := net.InterfaceByName(i.Name)

		if err != nil {
			continue
		}

		if IsVirtualInterface(found.Name) {
			continue
		}

		iface = found
		break
	}

	if iface == nil {
		return ""
	}

	client, err := arp.Dial(iface)

	if err != nil {
		return ""
	}

	defer client.Close()

	addr, err := netip.ParseAddr(ip)

	if err != nil {
		return ""
	}

	err = client.SetReadDeadline(
		time.Now().Add(
			300 * time.Millisecond,
		),
	)

	if err != nil {
		return ""
	}

	mac, err := client.Resolve(addr)

	if err != nil {
		return ""
	}

	return mac.String()
}
