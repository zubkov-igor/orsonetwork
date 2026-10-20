package wifi

type AccessPoint struct {
	BSSID      string `json:"bssid"`
	Channel    string `json:"channel"`
	Signal     string `json:"signal,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	SSID       string `json:"ssid"`
}

type Client struct {
	MAC    string `json:"mac"`
	BSSID  string `json:"bssid,omitempty"`
	Signal string `json:"signal,omitempty"`
}
