package scanner

import "testing"

func TestWailsTopology(t *testing.T) {

	s := New()

	topology := s.Topology()

	t.Log("Networks:", len(topology.Networks))
	t.Log("Nodes:", len(topology.Nodes))
	t.Log("Links:", len(topology.Links))

	for _, n := range topology.Nodes {
		t.Log(
			n.Type,
			n.IP,
			n.Vendor,
		)
	}
}