package scanner

import (
	"context"
	"net"
	"strings"
	"time"

	"OrsoNetwork/internal/logger"
)

func LookupReverseDNS(ip string) string {

	const dnsTimeout = 500 * time.Millisecond

	ctx, cancel := context.WithTimeout(
		context.Background(),
		dnsTimeout,
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

	logger.Log.Println(
    "REVERSE DNS:",
    ip,
    names,
)

	if len(names) == 0 {
		return ""
	}

	return strings.TrimSuffix(
		names[0],
		".",
	)
}
