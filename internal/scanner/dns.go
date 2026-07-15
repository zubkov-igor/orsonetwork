package scanner

import (
	"net"
	"strings"
)

func LookupHostname(ip string) string {

	names, err := net.LookupAddr(ip)

	if err != nil {
		return ""
	}

	if len(names) == 0 {
		return ""
	}

	return strings.TrimSuffix(
		names[0],
		".",
	)
}