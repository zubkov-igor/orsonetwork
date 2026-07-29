package scanner

import (
    "strings"
)

func ParseNetBIOSResponse(
    data []byte,
) NetBIOSResult {

    result := NetBIOSResult{}

    text := string(data)


    // имя компьютера
    lines := strings.Split(
        text,
        "\n",
    )


    for _, line := range lines {

        line = strings.TrimSpace(line)


        if result.Name == "" &&
            strings.Contains(line, "<00>") &&
            !strings.Contains(line, "<GROUP>") {

            fields := strings.Fields(line)

            if len(fields) > 0 {
                result.Name = fields[0]
            }
        }


        if strings.Contains(line, "MAC Address") {

            parts := strings.Split(
                line,
                "=",
            )

            if len(parts) == 2 {

                mac := strings.TrimSpace(
                    parts[1],
                )

                if validNetBIOSMAC(mac) {
                    result.MAC = mac
                }
            }
        }
    }


    return result
}


func validNetBIOSMAC(mac string) bool {

    if mac == "" {
        return false
    }

    if mac == "00-00-00-00-00-00" {
        return false
    }

    return true
}