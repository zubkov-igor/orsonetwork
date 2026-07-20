//go:build windows

package scanner

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"OrsoNetwork/internal/models"
)

func getGateways() []models.Gateway {

	var result []models.Gateway

	println("EXECUTING ROUTE")

	cmd := exec.Command(
		"route",
		"print",
		"-4",
	)

	println("COMMAND CREATED")

	output, err := cmd.Output()

	println("COMMAND FINISHED")

	if err != nil {
		println(
			"COMMAND ERROR:",
			err.Error(),
		)

		return result
	}

	exe, err := os.Executable()

	if err == nil {

		dir := filepath.Dir(exe)

		logFile := filepath.Join(
			dir,
			"debug.txt",
		)

		os.WriteFile(
			logFile,
			output,
			0644,
		)
	}

	scanner := bufio.NewScanner(
		strings.NewReader(
			string(output),
		),
	)

	for scanner.Scan() {

		fields := strings.Fields(
			scanner.Text(),
		)

		if len(fields) < 5 {
			continue
		}

		if fields[0] != "0.0.0.0" {
			continue
		}

		if fields[1] != "0.0.0.0" {
			continue
		}

		println(
			"GATEWAY FOUND:",
			fields[2],
			fields[3],
		)

		result = append(
			result,
			models.Gateway{
				IP:        fields[2],
				Interface: fields[3],
			},
		)
	}

	return result
}
