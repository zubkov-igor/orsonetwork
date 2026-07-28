package scanner

func ProbeMDNS(
    ip string,
) UDPProbeResult {

    return UDPProbeResult{
        Found: false,
        Info:  "",
    }
}