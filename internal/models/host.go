package models

import "time"

type Host struct {

	IP string
	MAC string
	Hostname string
	Vendor string
	Online bool
	RTT time.Duration
}
