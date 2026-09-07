package scanner

import (
	"embed"
	"encoding/json"
	"strings"
)

//go:embed data/mobile.json
var deviceModelsFS embed.FS

type DeviceModel struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

func LookupDeviceModel(code string) (DeviceModel, bool) {
	data, err := deviceModelsFS.ReadFile("data/mobile.json")
	if err != nil {
		return DeviceModel{}, false
	}

	var models map[string]DeviceModel

	if err := json.Unmarshal(data, &models); err != nil {
		return DeviceModel{}, false
	}

	model, ok := models[code]

	return model, ok
}

func FindDeviceModelFromHostname(hostname string) (DeviceModel, bool) {
	hostname = strings.TrimSpace(hostname)
	hostname = strings.TrimSuffix(hostname, ".")

	if hostname == "" {
		return DeviceModel{}, false
	}

	// Берём только первую DNS label.
	code := strings.SplitN(hostname, ".", 2)[0]

	// Коды в базе записаны с разным регистром.
	code = strings.ToUpper(code)

	return LookupDeviceModel(code)
}