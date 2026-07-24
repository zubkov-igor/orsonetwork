package scanner

import (
        "strings"
)

func ParseNMBLookup(data []byte) NetBIOSResult {

        result := NetBIOSResult{}

        lines := strings.Split(
                string(data),
                "\n",
        )

for _, line := range lines {

        line = strings.TrimSpace(line)

        if result.Name == "" &&
                strings.Contains(line, "<00>") &&
                !strings.Contains(line, "<GROUP>") {

                parts := strings.Fields(line)

                if len(parts) > 0 {
                        result.Name = parts[0]
                }
        }


        if strings.Contains(line, "MAC Address") {

                parts := strings.Split(
                        line,
                        "=",
                )

                if len(parts) == 2 {

                        result.MAC = strings.TrimSpace(
                                parts[1],

                        )
                }
        }
}

    return result
}