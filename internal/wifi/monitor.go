package wifi

import (
	"fmt"
	"os/exec"
)

const (
	iwBin = "/usr/sbin/iw"
	ipBin = "/usr/sbin/ip"
)

func enableMonitor(iface string) error {
	if output, err := exec.Command(
		"nmcli",
		"device", "set", iface, "managed", "no",
	).CombinedOutput(); err != nil {
		return fmt.Errorf(
			"set interface unmanaged: %w: %s",
			err,
			string(output),
		)
	}

	if err := exec.Command(
		ipBin,
		"link", "set", iface, "down",
	).Run(); err != nil {
		return fmt.Errorf("bring interface down: %w", err)
	}

	if output, err := exec.Command(
		iwBin,
		"dev", iface, "set", "type", "monitor",
	).CombinedOutput(); err != nil {
		return fmt.Errorf(
			"set monitor mode: %w: %s",
			err,
			string(output),
		)
	}

	if err := exec.Command(
		ipBin,
		"link", "set", iface, "up",
	).Run(); err != nil {
		return fmt.Errorf("bring interface up: %w", err)
	}

	return nil
}

func disableMonitor(iface string) error {
	if err := exec.Command(
		ipBin,
		"link", "set", iface, "down",
	).Run(); err != nil {
		return fmt.Errorf("bring interface down: %w", err)
	}

	if output, err := exec.Command(
		iwBin,
		"dev", iface, "set", "type", "managed",
	).CombinedOutput(); err != nil {
		return fmt.Errorf(
			"set managed mode: %w: %s",
			err,
			string(output),
		)
	}

	if err := exec.Command(
		ipBin,
		"link", "set", iface, "up",
	).Run(); err != nil {
		return fmt.Errorf("bring interface up: %w", err)
	}

	if output, err := exec.Command(
		"nmcli",
		"device", "set", iface, "managed", "yes",
	).CombinedOutput(); err != nil {
		return fmt.Errorf(
			"set interface managed: %w: %s",
			err,
			string(output),
		)
	}

	return nil
}
