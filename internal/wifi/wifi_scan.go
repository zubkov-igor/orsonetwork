package wifi

import (
	"bytes"
	"os/exec"
	"strings"
)

type Network struct {
	SSID     string `json:"ssid"`
	BSSID    string `json:"bssid,omitempty"`
	Signal   string `json:"signal,omitempty"`
	Security string `json:"security,omitempty"`
	Channel  string `json:"channel,omitempty"`
	Freq     string `json:"freq,omitempty"`
	Rate     string `json:"rate,omitempty"`
	Mode     string `json:"mode,omitempty"`
}

func Scan() ([]Network, error) {
	// Сначала пробуем nmcli (предпочтительно)
	networks, err := scanNmcli()
	if err == nil && len(networks) > 0 {
		return networks, nil
	}

	// Fallback на iw
	return scanIw()
}

func scanNmcli() ([]Network, error) {
	cmd := exec.Command("nmcli", "-t", "-f",
		"SSID,BSSID,SIGNAL,SECURITY,CHAN,FREQ,RATE,MODE",
		"device", "wifi", "list", "--rescan", "yes")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var networks []Network
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		// nmcli -t разделяет поля через ':'
		// BSSID содержит ':', поэтому парсим аккуратно
		parts := splitNmcliLine(line, 8)
		if len(parts) < 8 {
			continue
		}

		ssid := parts[0]
		if ssid == "" || ssid == "--" {
			continue
		}

		networks = append(networks, Network{
			SSID:     ssid,
			BSSID:    parts[1],
			Signal:   parts[2] + "%",
			Security: parts[3],
			Channel:  parts[4],
			Freq:     parts[5],
			Rate:     parts[6],
			Mode:     parts[7],
		})
	}

	return networks, nil
}

// splitNmcliLine корректно разбирает строку nmcli с учётом ':' в BSSID
func splitNmcliLine(line string, expected int) []string {
	parts := strings.Split(line, ":")
	if len(parts) <= expected {
		return parts
	}

	// BSSID занимает несколько частей (aa:bb:cc:dd:ee:ff)
	result := make([]string, 0, expected)
	result = append(result, parts[0]) // SSID

	// Собираем BSSID (следующие 6 частей)
	if len(parts) >= 7 {
		bssid := strings.Join(parts[1:7], ":")
		result = append(result, bssid)
		result = append(result, parts[7:]...)
	} else {
		result = append(result, parts[1:]...)
	}

	return result
}

func scanIw() ([]Network, error) {
	// Определяем беспроводной интерфейс
	iface, err := findWirelessInterface()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("iw", "dev", iface, "scan")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parseIwScan(string(out)), nil
}

func findWirelessInterface() (string, error) {
	out, err := exec.Command("iw", "dev").Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Interface ") {
			return strings.TrimPrefix(line, "Interface "), nil
		}
	}
	return "", exec.ErrNotFound
}

func parseIwScan(output string) []Network {
	var networks []Network
	var current Network

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "BSS "):
			if current.SSID != "" {
				networks = append(networks, current)
			}
			current = Network{}
			// BSS aa:bb:cc:dd:ee:ff
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				current.BSSID = strings.TrimSuffix(fields[1], "(")
			}

		case strings.HasPrefix(line, "SSID:"):
			current.SSID = strings.TrimSpace(strings.TrimPrefix(line, "SSID:"))

		case strings.HasPrefix(line, "signal:"):
			// signal: -45.00 dBm
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				current.Signal = parts[1] + " " + parts[2]
			}

		case strings.HasPrefix(line, "freq:"):
			current.Freq = strings.TrimSpace(strings.TrimPrefix(line, "freq:"))

		case strings.Contains(line, "primary channel:"):
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				current.Channel = strings.TrimSpace(parts[1])
			}
		}
	}

	if current.SSID != "" {
		networks = append(networks, current)
	}

	return networks
}
