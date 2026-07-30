package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)

func IdentifyDevice(
	host models.Host,
) models.DeviceType {

	hostname := strings.ToLower(
		host.Hostname,
	)

	vendor := strings.ToLower(
		host.Vendor,
	)

	// Gateway
	if host.IP == "192.168.0.1" {
		return models.DeviceGateway
	}

	// Routers by vendor
	if strings.Contains(vendor, "eltex") ||
		strings.Contains(vendor, "tp-link") ||
		strings.Contains(vendor, "xiaomi") {

		return models.DeviceRouter
	}

	// Windows computers
	if strings.HasPrefix(hostname, "desktop") ||
		strings.Contains(hostname, "win") ||
		strings.Contains(hostname, "pc-") {

		return models.DeviceComputer
	}

	// Phones
	if strings.Contains(hostname, "iphone") ||
		strings.Contains(hostname, "android") {

		return models.DevicePhone
	}

	// SMB ports
	for _, port := range host.Ports {

		if port.Number == 445 ||
			port.Number == 139 {

			return models.DeviceComputer
		}
	}

	// mDNS mobile hints
	for _, udp := range host.UDPServices {

		if strings.Contains(
			strings.ToLower(udp.Service),
			"airplay",
		) {

			return models.DevicePhone
		}
	}

	return models.DeviceUnknown
}
