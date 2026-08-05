package scanner

var httpFingerprintRules = []HTTPFingerprintRule{

	{
		Needle:      "boardtype",
		Fingerprint: "router-ui",
	},
	{
		Needle:      "softwareversion",
		Fingerprint: "router-firmware",
	},
	{
		Needle:      "firmwareversion",
		Fingerprint: "router-firmware",
	},
	{
		Needle:      "signin",
		Fingerprint: "router-login",
	},

	{
    Needle: "lighttpd",
    Fingerprint: "embedded-web-server",
	},
}