package scanner

import "testing"

func TestLookupHostname(t *testing.T) {

	hostname := LookupHostname("192.168.0.1")

	t.Log("Hostname:", hostname)
}
