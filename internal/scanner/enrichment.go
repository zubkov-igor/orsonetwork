package scanner

import (
    "time"

    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func EnrichHosts(
    hosts []models.Host,
    iface models.Interface,
) []models.Host {

    logger.Log.Println(
        "HOSTS BEFORE ENRICHMENT:",
        len(hosts),
    )

    enrichmentStart := time.Now()

    // =========================
    // mDNS
    // =========================

    stepStart := time.Now()

    hosts = EnrichMDNS(
        hosts,
    )

    logger.Log.Println(
        "ENRICH TIMING MDNS:",
        time.Since(stepStart),
    )

    // =========================
    // UDP
    // =========================

    stepStart = time.Now()

    hosts = EnrichUDP(
        hosts,
        iface,
    )

    logger.Log.Println(
        "ENRICH TIMING UDP:",
        time.Since(stepStart),
    )

    // =========================
    // TCP Ports
    // =========================

    stepStart = time.Now()

    hosts = EnrichPorts(
        hosts,
    )

    logger.Log.Println(
        "ENRICH TIMING PORTS:",
        time.Since(stepStart),
    )


    for i := range hosts {

    UpdateHostStatus(
        &hosts[i],
    )
}

    // =========================
    // Identification
    // =========================

    stepStart = time.Now()

    for i := range hosts {

        identification := IdentifyDevice(
            hosts[i],
        )

        hosts[i].Type = identification.Type

        hosts[i].Confidence =
            CalculateConfidence(
                hosts[i],
            )

        logger.Log.Println(
            "IDENTIFICATION:",
            hosts[i].IP,
            "TYPE:",
            hosts[i].Type,
            "CONFIDENCE:",
            hosts[i].Confidence,
        )
    }

    logger.Log.Println(
        "ENRICH TIMING IDENTIFICATION:",
        time.Since(stepStart),
    )

    logger.Log.Println(
        "ENRICH TIMING TOTAL:",
        time.Since(enrichmentStart),
    )

    logger.Log.Println(
        "HOSTS BEFORE IDENTIFICATION:",
        len(hosts),
    )

    return hosts
}

func UpdateHostStatus(
    host *models.Host,
) {

    // ICMP

    if host.Online {

        logger.Log.Println(
            "HOST STATUS:",
            host.IP,
            "ONLINE",
            "REASON: ICMP",
        )

        return
    }

    // ARP

    if host.MAC != "" {

        host.Online = true

        logger.Log.Println(
            "HOST STATUS:",
            host.IP,
            "ONLINE",
            "REASON: ARP",
        )

        return
    }

    // TCP

    if len(host.Ports) > 0 {

        host.Online = true

        logger.Log.Println(
            "HOST STATUS:",
            host.IP,
            "ONLINE",
            "REASON: TCP PORT",
        )

        return
    }

    // UDP

    if len(host.UDPServices) > 0 {

        host.Online = true

        logger.Log.Println(
            "HOST STATUS:",
            host.IP,
            "ONLINE",
            "REASON: UDP SERVICE",
        )

        return
    }

    // HTTP

    if len(host.HTTP) > 0 {

        host.Online = true

        logger.Log.Println(
            "HOST STATUS:",
            host.IP,
            "ONLINE",
            "REASON: HTTP",
        )

        return
    }

    host.Online = false

    logger.Log.Println(
        "HOST STATUS:",
        host.IP,
        "OFFLINE",
        "REASON: NO RESPONSE",
    )
}