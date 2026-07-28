package scanner

import (
	"net"
	"time"
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
			Found:false,
		}
	}


	defer conn.Close()



	_, err = conn.Write(
		snmpRequest,
	)


	if err != nil {

		return UDPProbeResult{
			Found:false,
		}
	}



	buffer := make(
		[]byte,
		2048,
	)



	conn.SetReadDeadline(
		time.Now().Add(
			2*time.Second,
		),
	)



	n, err := conn.Read(
		buffer,
	)


	if err != nil {

		return UDPProbeResult{
			Found:false,
		}
	}



	response := buffer[:n]



	if len(response) < 2 {

		return UDPProbeResult{
			Found:false,
		}
	}



	if response[0] != 0x30 {

		return UDPProbeResult{
			Found:false,
		}
	}



	for i := 0; i < len(response); i++ {


		if response[i] == 0xA2 {


			info := string(response)


			return UDPProbeResult{
				Found:true,
				Info:info,
			}
		}
	}



	return UDPProbeResult{
		Found:false,
	}
}