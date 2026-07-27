package scanner

import (
	"testing"
)

func TestHostsFromCIDR(t *testing.T) {

	hosts := HostsFromCIDR(
		"192.168.0.0/29",
	)

	expected := 6

	if len(hosts) != expected {

		t.Fatalf(
			"expected %d hosts, got %d",
			expected,
			len(hosts),
		)
	}

	t.Log(hosts)
}
