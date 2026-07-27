package scanner

// CommonPorts — небольшой набор наиболее информативных портов
// для идентификации устройств.
var CommonPorts = []int{
    22,   // SSH
    53,   // DNS
    80,   // HTTP
    139,  // NetBIOS
    443,  // HTTPS
    445,  // SMB
    515,  // LPD Printer
    554,  // RTSP Camera
    631,  // IPP Printer
    9100, // RAW Printer
    3389, // RDP
}