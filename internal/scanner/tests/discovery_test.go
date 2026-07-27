package scanner

import (
	"testing"
)

func TestDiscoverHosts(t *testing.T) {

	hosts := DiscoverHosts(
		"192.168.0.0/30",
	)

	for _, h := range hosts {

		t.Log(
			h.IP,
			h.MAC,
			h.Vendor,
			h.Online,
			h.RTT,
		)
	}

}
