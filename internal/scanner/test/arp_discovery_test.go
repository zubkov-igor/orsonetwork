package scanner

import (
	"net"
	"net/netip"
	"testing"

	"github.com/mdlayher/arp"
)

func TestARPRequest(t *testing.T) {

	iface, err := net.InterfaceByName("enp4s0")
	if err != nil {
		t.Fatal(err)
	}

	client, err := arp.Dial(iface)
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	mac, err := client.Resolve(
		netip.MustParseAddr("192.168.0.1"),
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(
		"MAC:",
		mac,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(
		"MAC:",
		mac,
	)

}

func TestARPDiscovery(t *testing.T) {

	hosts := ARPDiscovery(
		"192.168.0.0/24",
	)

	for _, h := range hosts {

		t.Log(
			h.IP,
			h.MAC,
			h.Online,
		)

	}

	if len(hosts) == 0 {
		t.Error("no hosts found")
	}
}
