package scanner

import (
    "net"
    "net/netip"
    "time"

    "OrsoNetwork/internal/models"
    "OrsoNetwork/internal/logger"

    "github.com/mdlayher/arp"
)


func ARPDiscovery(
    hosts []models.Host,
) []models.Host {

    logger.Log.Println(
        "ARP DISCOVERY START",
    )

    var arpHosts []models.Host


    interfaces := GetInterfaces()

    for _, i := range interfaces {

    logger.Log.Println(
        "AVAILABLE INTERFACE:",
        i.Name,
    )
}


    if len(interfaces) == 0 {
        return arpHosts
    }


var iface *net.Interface


for _, i := range interfaces {

    if i.Name == "" {
        continue
    }


    found, err := net.InterfaceByName(
        i.Name,
    )

    if err != nil {
        continue
    }


    if IsVirtualInterface(found.Name) {

        logger.Log.Println(
            "SKIP INTERFACE:",
            found.Name,
        )

        continue
    }


    iface = found

    logger.Log.Println(
        "ARP INTERFACE:",
        iface.Name,
        iface.HardwareAddr.String(),
    )

    break
}


    if iface == nil {
        logger.Log.Println(
            "NO ARP INTERFACE",
        )

        return arpHosts
    }



    client, err := arp.Dial(
        iface,
    )


    if err != nil {

        logger.Log.Println(
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


       logger.Log.Println(
    "ARP REQUEST:",
    host.IP,
)

        err = client.SetReadDeadline(
            time.Now().Add(
                1 * time.Second,
            ),
        )


        if err != nil {
            continue
        }


        mac, err := client.Resolve(
            addr,
        )


        if err != nil {

            logger.Log.Println(
    "ARP RESOLVE ERROR:",
    host.IP,
    err.Error(),
)

            arpHosts = append(
                arpHosts,
                host,
            )

            continue
        }



        arpHost := host

        arpHost.MAC = mac.String()



     logger.Log.Println(
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
