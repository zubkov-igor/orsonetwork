package scanner

import (
	"OrsoNetwork/internal/models"
)

func ProbeMDNS(
	ip string,
	iface models.Interface,
) UDPProbeResult {

	return UDPProbeResult{
		Found: false,
		Info:  "",
	}
}
