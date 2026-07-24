package scanner

import (
	"context"
	"os/exec"
	"time"
)

func runNMBLookup(ip string) ([]byte, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"nmblookup",
		"-A",
		ip,
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return output, err
	}

	return output, nil
}