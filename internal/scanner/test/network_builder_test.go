package scanner

import "testing"

func TestBuildNetwork(t *testing.T) {

	ifaces := GetInterfaces()
	gateways := GetGateways()

	for _, iface := range ifaces {

		gw := GatewayForInterface(
			iface.Name,
			gateways,
		)

		if gw == nil {
			continue
		}

		network := BuildNetwork(
			iface,
			gw,
		)

		t.Log(
			network.Interface,
			network.CIDR,
			network.Gateway,
		)
	}
}
