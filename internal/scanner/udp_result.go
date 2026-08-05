package scanner

type UDPProbeResult struct {
    Found bool
    Info string
    Raw  []byte
}