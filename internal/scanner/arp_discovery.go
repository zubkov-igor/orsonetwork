package scanner

import (
    "net"
    "net/netip"
    "time"

    "OrsoNetwork/internal/models"

    "github.com/mdlayher/arp"
)


func ARPDiscovery(cidr string) []models.Host {

    var hosts []models.Host


    interfaces := GetInterfaces()

    if len(interfaces) == 0 {
        return hosts
    }


    iface, err := net.InterfaceByName(
        interfaces[0].Name,
    )

    if err != nil {
        return hosts
    }


    client, err := arp.Dial(iface)

    if err != nil {
        return hosts
    }

    defer client.Close()



    // время ожидания ARP ответов
    err = client.SetReadDeadline(
        time.Now().Add(
            2 * time.Second,
        ),
    )

    if err != nil {
        return hosts
    }



    ips := HostsFromCIDR(cidr)



    // отправляем ARP запросы
    for _, ip := range ips {


        addr, err := netip.ParseAddr(ip)

        if err != nil {
            continue
        }


        err = client.Request(addr)

        if err != nil {
            continue
        }

    }



    // читаем ответы
    for {


        packet, _, err := client.Read()


        if err != nil {
            break
        }



        hosts = append(
            hosts,
            models.Host{
                IP: packet.SenderIP.String(),
                MAC: packet.SenderHardwareAddr.String(),
                Online: true,
            },
        )

    }



    return hosts
}