package scanner

var PortServices = map[int]string{

	22: "ssh",
	23: "telnet",

	53: "dns",

	80:  "http",
	443: "https",

	135: "msrpc",
	139: "netbios",
	445: "smb",

	515:  "lpd-printer",
	631:  "ipp-printer",
	9100: "raw-printer",

	554: "rtsp",

	1900: "ssdp-upnp",

	3389: "rdp",

	5000: "synology",

	8000: "camera-http",

	8008: "google-cast",
	8009: "chromecast",

	8080: "http-alt",
	8443: "https-alt",
}

func DetectPortService(port int) string {

	service, ok := PortServices[port]

	if ok {
		return service
	}

	return "unknown"
}
