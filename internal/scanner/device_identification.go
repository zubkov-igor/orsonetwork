package scanner

import (
	"strings"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func IdentifyDevice(
	host models.Host,
	evidence []models.FingerprintEvidence,
) models.DeviceIdentification {

	logger.Debug(
		"BEFORE IDENTIFICATION:",
		host.IP,
		"HOSTNAME:", host.Hostname,
		"MAC:", host.MAC,
		"PORTS:", len(host.Ports),
		"UDP:", len(host.UDPServices),
		"MDNS:", len(host.MDNS),
	)

	hostname := strings.ToLower(host.Hostname)
vendor := strings.ToLower(host.Vendor)

scores := make(
    map[models.DeviceType]int,
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
		scores[models.DeviceRouter] += 80
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
		scores[models.DeviceRouter] += 50
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
		scores[models.DeviceCamera] += 80
	}

	// =========================
	// Printer
	// =========================

	if containsAny(
		hostname,
		"printer",
		"print",
	) {
		scores[models.DevicePrinter] += 80
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
		scores[models.DeviceNAS] += 80
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
		scores[models.DeviceComputer] += 70
	}

	// =========================
	// Ports
	// =========================

	for _, port := range host.Ports {

		switch port.Number {

		case 22:
			scores[models.DeviceServer] += 20

		case 3389:
			scores[models.DeviceComputer] += 20

		case 445:
			scores[models.DeviceComputer] += 15
			scores[models.DeviceNAS] += 15
		}
	}

	// =========================
	// HTTP
	// =========================

	for _, httpInfo := range host.HTTP {

		server := strings.ToLower(
			httpInfo.Server,
		)

		if strings.Contains(
			server,
			"apache",
		) {
			scores[models.DeviceComputer] += 10
		}

		if strings.Contains(
			server,
			"nginx",
		) {
			scores[models.DeviceComputer] += 10
		}

		if strings.Contains(
			server,
			"lighttpd",
		) {
			scores[models.DeviceRouter] += 20
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
		scores[models.DeviceIoT] += 50
	}

	// =========================
	// Find best match
	// =========================

	var bestType models.DeviceType
	bestScore := 0

	for deviceType, score := range scores {

		if score > bestScore {

			bestType = deviceType
			bestScore = score
		}
	}

	if bestScore == 0 {

		return models.DeviceIdentification{
			Type:       models.DeviceUnknown,
			Confidence: 0,
		}
	}

	if bestScore > 100 {
		bestScore = 100
	}

	return models.DeviceIdentification{
		Type:       bestType,
		Confidence: bestScore,
	}
}

func IdentifyOS(host models.Host) string {

	logger.Debug(
		"OS FINGERPRINT:",
		host.IP,
		"PORTS:", len(host.Ports),
		"UDP:", len(host.UDPServices),
		"MDNS:", len(host.MDNS),
	)

	for _, port := range host.Ports {
		logger.Debug(
			"OS PORT:",
			host.IP,
			port.Number,
			port.Protocol,
			port.Service,
		)
	}

	for _, service := range host.UDPServices {
		logger.Debug(
			"OS UDP:",
			host.IP,
			service.Port,
			service.Service,
		)
	}

	for _, service := range host.MDNS {
		logger.Debug(
			"OS MDNS:",
			host.IP,
			service.Service,
		)
	}

	score := 0

	// =========================
	// TCP ports
	// =========================

	for _, port := range host.Ports {

		switch port.Number {

		case 135:
			// Microsoft RPC — слабый признак Windows
			score += 10

		case 139:
			// NetBIOS Session Service — сильный признак Windows
			score += 25

		case 445:
			// SMB — сильный признак Windows
			score += 30

		case 3389:
			// RDP — сильный дополнительный признак Windows
			score += 25
		}
	}

	// =========================
	// UDP services
	// =========================

	for _, service := range host.UDPServices {

		if service.Port == 137 &&
			strings.EqualFold(
				service.Service,
				"NetBIOS",
			) {

			score += 25
		}
	}

	// =========================
	// mDNS services
	// =========================

	for _, service := range host.MDNS {

		if strings.Contains(
			strings.ToLower(service.Service),
			"_dosvc._tcp",
		) {
			score += 30
		}
	}

	// =========================
	// Result
	// =========================

	// Routers can expose SMB/NetBIOS services,
	// so these ports alone are not enough to identify Windows.
	if host.Type == models.DeviceRouter {
		return "unknown"
	}

	if score >= 40 {
		return "Windows"
	}

	return "unknown"
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

func hasEvidence(
    evidence []models.FingerprintEvidence,
    value string,
) bool {

    for _, item := range evidence {

        if strings.EqualFold(
            item.Value,
            value,
        ) {
            return true
        }
    }

    return false
}
