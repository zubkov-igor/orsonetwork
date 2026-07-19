package scanner

import "testing"

func TestGetGateways(t *testing.T) {

	gateways := GetGateways()


	for _, gw := range gateways {

		t.Log(
			gw.Interface,
			gw.IP,
		)

	}


	if len(gateways) == 0 {
		t.Error("no gateways found")
	}
}