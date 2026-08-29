package models

type ScannerConfig struct {

    EnableICMP bool
    EnableARP bool

    EnableReverseDNS bool
    EnableNetBIOS bool

    EnableMDNS bool
    EnableSSDP bool

    EnableSNMP bool

    EnableTCP bool
    EnableUDP bool

    Workers int
}