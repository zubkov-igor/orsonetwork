package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)

func IdentifyDevice(
	host models.Host,
) models.DeviceType {

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
		return models.DeviceRouter
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
		return models.DeviceRouter
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
		return models.DeviceCamera
	}

	// =========================
	// Printer
	// =========================

	if containsAny(
		hostname,
		"printer",
		"print",
	) {
		return models.DevicePrinter
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
		return models.DeviceNAS
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
		return models.DeviceComputer
	}

	// =========================
	// Server
	// =========================

	for _, port := range host.Ports {

		switch port.Number {

		case 22:
			return models.DeviceServer

		case 3389:
			return models.DeviceComputer

		case 80, 443:

			if containsAny(
				hostname,
				"server",
			) {
				return models.DeviceServer
			}
		}
	}

	// =========================
	// IoT
	// =========================

	if containsAny(
		vendor,
		"shenzhen",
		"esp",
		"tuya",
		"sonoff",
	) {
		return models.DeviceIoT
	}

	return models.DeviceUnknown
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
