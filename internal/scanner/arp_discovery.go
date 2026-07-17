package scanner

import (
    "net"
    "net/netip"
  
    "OrsoNetwork/internal/models"

    "github.com/mdlayher/arp"
)

func ARPDiscovery(
    hosts []models.Host,
) []models.Host {

    var arpHosts []models.Host

    interfaces := GetInterfaces()

    if len(interfaces) == 0 {
        return arpHosts
    }

    iface, err := net.InterfaceByName(
        interfaces[0].Name,
    )

    if err != nil {
        return arpHosts
    }

    client, err := arp.Dial(iface)

    if err != nil {
        println(
            "ARP DIAL ERROR:",
            err.Error(),
        )

        return arpHosts
    }

    defer client.Close()

    for _, host := range hosts {

        addr, err := netip.ParseAddr(
            host.IP,
        )

        if err != nil {
            continue
        }

        println(
            "ARP REQUEST:",
            host.IP,
        )

        mac, err := client.Resolve(addr)

        if err != nil {

            println(
                "ARP RESOLVE ERROR:",
                host.IP,
                err.Error(),
            )

            continue
        }

        arpHost := models.Host{
            IP:     host.IP,
            MAC:    mac.String(),
            Online: true,
        }

        println(
            "ARP HOST:",
            arpHost.IP,
            arpHost.MAC,
        )

        arpHosts = append(
            arpHosts,
            arpHost,
        )
    }

    return arpHosts
}