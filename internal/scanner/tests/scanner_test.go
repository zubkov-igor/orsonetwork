package scanner

import "testing"

func TestScannerScan(t *testing.T) {

	s := New()

	networks := s.Scan()

	if len(networks) == 0 {
		t.Fatal("no networks found")
	}

	for _, network := range networks {

		t.Log(
			"Interface:",
			network.Interface,
		)

		t.Log(
			"CIDR:",
			network.CIDR,
		)

		t.Log(
			"Gateway:",
			network.Gateway,
		)

		t.Log(
			"Hosts:",
			len(network.Hosts),
		)

		if network.CIDR == "" {
			t.Error("empty CIDR")
		}

		if network.Gateway == "" {
			t.Error("empty gateway")
		}
	}
}
