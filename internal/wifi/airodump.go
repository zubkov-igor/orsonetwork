package wifi

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func AirodumpScan(
	ctx context.Context,
	iface string,
	channel int,
	duration time.Duration,
) ([]AccessPoint, []Client, error) {
	if iface == "" {
		return nil, nil, fmt.Errorf("wireless interface is empty")
	}

	if channel <= 0 {
		return nil, nil, fmt.Errorf("invalid wireless channel: %d", channel)
	}

	dir, err := os.MkdirTemp("", "orsonetwork-airodump-*")
	if err != nil {
		return nil, nil, err
	}
	defer os.RemoveAll(dir)

	prefix := filepath.Join(dir, "scan")

	scanCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	cmd := exec.CommandContext(
		scanCtx,
		airodumpBin,
		"--output-format", "csv",
		"--channel", strconv.Itoa(channel),
		"--write", prefix,
		"--write-interval", "1",
		iface,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf(
			"start airodump-ng: %w",
			err,
		)
	}

	if err := cmd.Wait(); err != nil {
		if !errors.Is(scanCtx.Err(), context.Canceled) &&
			!errors.Is(scanCtx.Err(), context.DeadlineExceeded) {
			return nil, nil, fmt.Errorf(
				"airodump-ng failed: %w",
				err,
			)
		}
	}

	csvPath := prefix + "-01.csv"

	data, err := os.ReadFile(csvPath)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"read airodump csv: %w",
			err,
		)
	}

	accessPoints, clients, err := parseAirodumpCSV(string(data))
	if err != nil {
		return nil, nil, fmt.Errorf(
			"parse airodump csv: %w",
			err,
		)
	}

	return accessPoints, clients, nil
}
