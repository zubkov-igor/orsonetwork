package scanner

import (
	"testing"

	"OrsoNetwork/internal/models"
)

func TestGatewayForInterface(t *testing.T) {

	gateways := []models.Gateway{
		{
			Interface: "enp4s0",
			IP:        "192.168.0.1",
		},
		{
			Interface: "virbr0",
			IP:        "192.168.122.1",
		},
	}

	gw := GatewayForInterface(
		"enp4s0",
		gateways,
	)

	if gw == nil {
		t.Fatal("gateway not found")
	}

	t.Log(
		gw.Interface,
		gw.IP,
	)
}
