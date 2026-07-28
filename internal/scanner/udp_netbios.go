package scanner

func ProbeNetBIOS(
    ip string,
) UDPProbeResult {

    return UDPProbeResult{
        Found:false,
        Info:"",
    }
}