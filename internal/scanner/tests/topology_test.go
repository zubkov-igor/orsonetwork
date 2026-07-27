package scanner

import "testing"

func TestBuildTopology(t *testing.T) {

	s := New()

	topology := s.Topology()

	t.Log("Networks:", len(topology.Networks))
	t.Log("Nodes:", len(topology.Nodes))
	t.Log("Links:", len(topology.Links))

	for _, network := range topology.Networks {

		t.Log(
			network.Interface,
			network.CIDR,
			network.Gateway,
			len(network.Hosts),
		)
	}

	for _, node := range topology.Nodes {

		t.Log(
			"NODE:",
			node.Type,
			node.ID,
			node.IP,
			node.MAC,
			node.Vendor,
		)
	}

	for _, link := range topology.Links {

		t.Log(
			"LINK:",
			link.From,
			"->",
			link.To,
			link.Type,
		)
	}
}
