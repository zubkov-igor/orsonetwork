package scanner

import (
	"time"

	"OrsoNetwork/internal/models"

	ping "github.com/go-ping/ping"
)


func PingHost(
	ip string,
	timeout time.Duration,
) models.Host {


	host := models.Host{
		IP: ip,
	}

	pinger, err := ping.NewPinger(ip)

	if err != nil {
		return host
	}


	pinger.Count = 1

	pinger.Timeout = timeout


	// ICMP требует root
	pinger.SetPrivileged(false)


	err = pinger.Run()

	if err != nil {

		host.Online = false

		return host
	}


	stats := pinger.Statistics()


	if stats.PacketsRecv == 0 {

		host.Online = false

		return host
	}


	host.Online = true

	host.RTT = stats.AvgRtt

	return host
}

