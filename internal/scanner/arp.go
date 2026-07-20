package scanner

import (
	"os/exec"
	"strings"
)

type ARPTable map[string]string

func (a ARPTable) Lookup(ip string) string {

	return a[ip]

}

func GetARPTable() ARPTable {

	table := make(ARPTable)

	cmd := exec.Command(
		"ip",
		"neigh",
	)

	output, err := cmd.Output()

	if err != nil {
		return table
	}

	lines := strings.Split(
		string(output),
		"\n",
	)

	for _, line := range lines {

		fields := strings.Fields(line)

		if len(fields) < 5 {
			continue
		}

		ip := fields[0]

		for i, field := range fields {

			if field == "lladdr" && i+1 < len(fields) {

				mac := fields[i+1]

				table[ip] = mac

				break
			}
		}
	}

	return table
}
