package wifi

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	iwBin       = "/usr/sbin/iw"
	ipBin       = "/usr/sbin/ip"
	airodumpBin = "/usr/sbin/airodump-ng"
)

var command = exec.Command

func createMonitorInterface(iface string) (string, error) {
	const monitorIface = "mon0"

	if output, err := exec.Command(
		iwBin,
		"dev", iface,
		"interface", "add",
		monitorIface,
		"type", "monitor",
	).CombinedOutput(); err != nil {
		return "", fmt.Errorf(
			"create monitor interface: %w: %s",
			err,
			strings.TrimSpace(string(output)),
		)
	}

	if output, err := exec.Command(
		ipBin,
		"link", "set", monitorIface, "up",
	).CombinedOutput(); err != nil {
		_ = exec.Command(
			iwBin,
			"dev", monitorIface,
			"del",
		).Run()

		return "", fmt.Errorf(
			"bring monitor interface up: %w: %s",
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return monitorIface, nil
}

func deleteMonitorInterface(iface string) error {
	if iface == "" {
		return nil
	}

	if output, err := exec.Command(
		iwBin, "dev", iface, "del",
	).CombinedOutput(); err != nil {
		return fmt.Errorf(
			"delete monitor interface: %w: %s",
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return nil
}

