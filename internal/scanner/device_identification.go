package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)

func IdentifyDevice(
	host models.Host,
) models.DeviceIdentification {

	hostname := strings.ToLower(
		host.Hostname,
	)

	vendor := strings.ToLower(
		host.Vendor,
	)

	// =========================
	// Router / Gateway
	// =========================

	if containsAny(
		hostname,
		"router",
		"gateway",
		"mikrotik",
		"openwrt",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceRouter,
			Confidence: 80,
		}
	}

	if containsAny(
		vendor,
		"d-link",
		"eltex",
		"mikrotik",
		"ubiquiti",
		"cisco",
		"netgear",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceRouter,
			Confidence: 50,
		}
	}

	// =========================
	// Camera
	// =========================

	if containsAny(
		hostname,
		"camera",
		"cam",
		"ipc",
		"nvr",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceCamera,
			Confidence: 80,
		}
	}

	// =========================
	// Printer
	// =========================

	if containsAny(
		hostname,
		"printer",
		"print",
	) {
		return models.DeviceIdentification{
			Type:       models.DevicePrinter,
			Confidence: 80,
		}
	}

	// =========================
	// NAS
	// =========================

	if containsAny(
		hostname,
		"nas",
		"storage",
		"synology",
		"qnap",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceNAS,
			Confidence: 80,
		}
	}

	// =========================
	// Computer
	// =========================

	if containsAny(
		hostname,
		"desktop",
		"pc",
		"laptop",
		"computer",
		"workstation",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceComputer,
			Confidence: 70,
		}
	}

	// =========================
	// Server
	// =========================

	for _, port := range host.Ports {

		switch port.Number {

		case 22:
			return models.DeviceIdentification{
				Type:       models.DeviceServer,
				Confidence: 60,
			}

		case 3389:
			return models.DeviceIdentification{
				Type:       models.DeviceComputer,
				Confidence: 60,
			}

		case 80, 443:

			if containsAny(
				hostname,
				"server",
			) {
				return models.DeviceIdentification{
					Type:       models.DeviceServer,
					Confidence: 70,
				}
			}
		}
	}

	// =========================
	// IoT
	// =========================

	if containsAny(
		vendor,
		"esp",
		"tuya",
		"sonoff",
	) {
		return models.DeviceIdentification{
			Type:       models.DeviceIoT,
			Confidence: 50,
		}
	}

	return models.DeviceIdentification{
		Type:       models.DeviceUnknown,
		Confidence: 0,
	}
}

func containsAny(
	value string,
	items ...string,
) bool {

	for _, item := range items {

		if strings.Contains(
			value,
			item,
		) {
			return true
		}
	}

	return false
}
