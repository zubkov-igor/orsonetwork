package scanner

import "OrsoNetwork/internal/models"


func CalculateConfidence(
    host models.Host,
) int {


    score := 0


    for _, source := range host.Sources {


        switch source.Type {


        case models.DiscoveryARP:
            score += 20


        case models.DiscoveryReverseDNS:
            score += 20


        case models.DiscoveryNetBIOS:
            score += 30


        case models.DiscoveryMDNS:
            score += 25


        }
    }


    if host.MAC != "" {
        score += 10
    }


    if len(host.Ports) > 0 {
        score += 10
    }


    if score > 100 {
        score = 100
    }


    return score
}