package scanner

import (
	"context"
	"net"
	"strings"
	"time"
)

func LookupHostname(ip string) string {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		200*time.Millisecond,
	)

	defer cancel()

	resolver := net.DefaultResolver

	names, err := resolver.LookupAddr(
		ctx,
		ip,
	)

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
