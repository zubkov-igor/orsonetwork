package wifi

type AccessPoint struct {
	BSSID      string `json:"bssid"`
	Channel    string `json:"channel"`
	Signal     string `json:"signal,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	SSID       string `json:"ssid"`
	Vendor     string `json:"vendor,omitempty"`
}

type Client struct {
	MAC    string `json:"mac"`
	BSSID  string `json:"bssid,omitempty"`
	Signal string `json:"signal,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Brand  string `json:"brand,omitempty"`
	Model  string `json:"model,omitempty"`
}
