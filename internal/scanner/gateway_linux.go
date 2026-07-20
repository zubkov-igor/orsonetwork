//go:build linux

package scanner

import (
	"OrsoNetwork/internal/models"

	"github.com/vishvananda/netlink"
)

func getGateways() []models.Gateway {

	var result []models.Gateway

	routes, err := netlink.RouteList(
		nil,
		netlink.FAMILY_V4,
	)

	if err != nil {
		return result
	}

	for _, route := range routes {

		if route.Dst != nil &&
			!route.Dst.IP.IsUnspecified() {
			continue
		}

		if route.Gw == nil {
			continue
		}

		link, err := netlink.LinkByIndex(
			route.LinkIndex,
		)

		if err != nil {
			continue
		}

		result = append(
			result,
			models.Gateway{
				Interface: link.Attrs().Name,
				IP:        route.Gw.String(),
			},
		)
	}

	return result
}
