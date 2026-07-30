package scanner

// CommonPorts — информативные TCP порты
var CommonPorts = []int{

	// Web
	80,   // HTTP
	443,  // HTTPS
	8080, // HTTP alternate
	8443, // HTTPS alternate

	// Remote access
	22,   // SSH
	23,   // Telnet
	3389, // RDP

	// Windows
	135, // RPC
	139, // NetBIOS
	445, // SMB

	// DNS
	53,

	// Printers
	515,
	631,
	9100,

	// Cameras / DVR
	554,  // RTSP
	8000, // Hikvision
	8081,

	// Media
	1900, // SSDP/UPnP TCP rare
}
