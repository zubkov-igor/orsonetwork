package scanner

import (
	"testing"
	"time"
)

func TestPingHost(t *testing.T) {

	host := PingHost(
		"192.168.0.1",
		time.Second,
	)

	if !host.Online {
		t.Fatal("host is offline")
	}

	t.Log(
		"RTT:",
		host.RTT,
	)
}
