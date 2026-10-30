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
	code = strings.ToUpper(code)

	// Сначала пробуем точное совпадение.
	if model, ok := LookupDeviceModel(code); ok {
		return model, true
	}

	// Некоторые устройства добавляют дополнительное имя
	// после кода модели: например M2004J19C-Redmi9.
	parts := strings.Split(code, "-")

	for i := len(parts) - 1; i > 0; i-- {
		candidate := strings.Join(parts[:i], "-")

		if model, ok := LookupDeviceModel(candidate); ok {
			return model, true
		}
	}

	return DeviceModel{}, false
}
