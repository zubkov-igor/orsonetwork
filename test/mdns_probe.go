package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

func main() {

	fmt.Println("mDNS TEST START")

	iface, err := net.InterfaceByName("enp4s0")

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"INTERFACE:",
		iface.Name,
		"INDEX:",
		iface.Index,
	)

	conn, err := net.ListenUDP(
		"udp4",
		&net.UDPAddr{
			IP:   net.IPv4zero,
			Port: 5353,
		},
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println(
		"LISTEN:",
		conn.LocalAddr(),
	)

	packetConn := ipv4.NewPacketConn(conn)

	group := &net.UDPAddr{
		IP: net.IPv4(
			224,
			0,
			0,
			251,
		),
	}

	err = packetConn.JoinGroup(
		iface,
		group,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"JOINED:",
		group.IP,
		"interface:",
		iface.Name,
	)

	fmt.Println("WAITING...")

	buffer := make([]byte, 4096)

	for {

		err = conn.SetReadDeadline(
			time.Now().Add(
				30 * time.Second,
			),
		)

		if err != nil {
			panic(err)
		}

		n, addr, err := conn.ReadFromUDP(
			buffer,
		)

		if err != nil {

			fmt.Println(
				"READ ERROR:",
				err,
			)

			continue
		}

		fmt.Println(
			"PACKET:",
			n,
			"bytes from",
			addr,
		)

		fmt.Printf(
			"%x\n",
			buffer[:n],
		)
	}
}