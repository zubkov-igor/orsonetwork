package scanner

import (
	"log"
	"net"
	"strings"
	"time"
)


func DiscoverSSDP() {


	addr := &net.UDPAddr{
		IP: net.IPv4zero,
		Port: 0,
	}


	conn, err := net.ListenUDP(
		"udp",
		addr,
	)


	if err != nil {
		return
	}


	defer conn.Close()



	target := &net.UDPAddr{
		IP: net.ParseIP(
			"239.255.255.250",
		),
		Port:1900,
	}



	request :=
		"M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 2\r\n" +
		"ST: ssdp:all\r\n" +
		"\r\n"



	_, err = conn.WriteToUDP(
		[]byte(request),
		target,
	)


	if err != nil {
		return
	}



	buffer := make(
		[]byte,
		2048,
	)



	deadline := time.Now().Add(
		5*time.Second,
	)



	for {


		conn.SetReadDeadline(
			deadline,
		)



		n, addr, err := conn.ReadFromUDP(
			buffer,
		)


		if err != nil {
			break
		}



		response :=
			strings.ToLower(
				string(buffer[:n]),
			)


		log.Println(
			"SSDP RESPONSE:",
			addr.IP,
		)


		log.Println(
			response,
		)

	}

}