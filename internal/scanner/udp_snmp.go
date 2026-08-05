package scanner

import (
	"fmt"
	"net"
	"time"
	"bytes"

	"OrsoNetwork/internal/logger"
)

func ProbeSNMP(
	ip string,
) UDPProbeResult {

	addr := ip + ":161"

	conn, err := net.DialTimeout(
		"udp",
		addr,
		2*time.Second,
	)

	if err != nil {

		return UDPProbeResult{
			Found: false,
		}
	}

	defer conn.Close()

logger.Log.Println(
    "OID:",
    OIDSysDescr,
)

request := BuildSNMPRequest(
    OIDSysDescr,
)

logger.Log.Println(
	"SNMP REQUEST HEX:",
	fmt.Sprintf("% X", request),
)

	_, err = conn.Write(request)

	if err != nil {

		return UDPProbeResult{
			Found: false,
		}
	}

	buffer := make(
		[]byte,
		2048,
	)

	conn.SetReadDeadline(
		time.Now().Add(
			2 * time.Second,
		),
	)

n, err := conn.Read(
	buffer,
)

if err != nil {
	return UDPProbeResult{
		Found: false,
	}
}

response := buffer[:n]

logger.Log.Println(
	"SNMP RESPONSE SIZE:",
	len(response),
)

logger.Log.Println(
	"SNMP RESPONSE HEX:",
	fmt.Sprintf("% X", response),
)

logger.Log.Println(
	"SNMP RESPONSE STRING:",
	string(response),
)

if response[0] != 0x30 {

	return UDPProbeResult{
		Found: false,
	}
}


if bytes.Contains(response, []byte{0xA2}) {

	return UDPProbeResult{
		Found: true,
		Info: "SNMP response",
		Raw: response,
	}
}


return UDPProbeResult{
	Found: false,
}
}
