package scanner

import "testing"

func TestBuildTopology(t *testing.T) {

	s := New()

	topology := s.Topology()

	t.Log("Networks:", len(topology.Networks))
	t.Log("Links:", len(topology.Links))

	for _, network := range topology.Networks {

		t.Log(
			network.Interface,
			network.CIDR,
			network.Gateway,
			len(network.Hosts),
		)
	}

	for _, link := range topology.Links {

		t.Log(
			link.From,
			"->",
			link.To,
			link.Type,
		)
	}
}