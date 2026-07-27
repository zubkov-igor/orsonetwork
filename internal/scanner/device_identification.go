package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)

func IdentifyDevice(
	device models.Device,
) models.DeviceType {

	hostname := strings.ToLower(device.Hostname)

	switch {

	case strings.Contains(hostname, "iphone"):
		return models.DevicePhone

	case strings.HasPrefix(hostname, "desktop"):
		return models.DeviceComputer

	case strings.Contains(hostname, "router"):
		return models.DeviceRouter

	case device.IP == "192.168.0.1":
		return models.DeviceGateway

	default:
		return models.DeviceUnknown
	}
}