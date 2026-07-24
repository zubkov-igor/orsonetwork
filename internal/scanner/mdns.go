package scanner

import "github.com/grandcat/zeroconf"

func DiscoverMDNS() map[string]string {

    _ = zeroconf.Resolver{}

    return map[string]string{}
}