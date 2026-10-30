package wifi

import (
	"encoding/csv"
	"fmt"
	"strings"

	"OrsoNetwork/internal/scanner"
)

func parseAirodumpCSV(data string) ([]AccessPoint, []Client, error) {
	reader := csv.NewReader(strings.NewReader(data))
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	var accessPoints []AccessPoint
	var clients []Client

	section := ""

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) == 0 {
			continue
		}

		first := strings.TrimSpace(record[0])

		switch {
		case first == "BSSID":
			section = "access_points"
			continue

		case first == "Station MAC":
			section = "clients"
			continue
		}

		switch section {
		case "access_points":
			if len(record) < 15 {
				continue
			}

			bssid := strings.TrimSpace(record[0])

			accessPoints = append(accessPoints, AccessPoint{
				BSSID:      bssid,
				Channel:    strings.TrimSpace(record[3]),
				Signal:     strings.TrimSpace(record[8]),
				Encryption: strings.TrimSpace(record[5]),
				SSID:       strings.TrimSpace(record[13]),
				Vendor:     scanner.LookupVendor(bssid),
			})

		case "clients":
			if len(record) < 7 {
				continue
			}

			mac := strings.TrimSpace(record[0])

			clients = append(clients, Client{
				MAC:    mac,
				Signal: strings.TrimSpace(record[3]),
				BSSID:  strings.TrimSpace(record[5]),
				Vendor: scanner.LookupVendor(mac),
			})
		}
	}

	if section == "" {
		return nil, nil, fmt.Errorf("airodump csv contains no known sections")
	}

	return accessPoints, clients, nil
}
