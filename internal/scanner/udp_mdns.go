package scanner

import (
	"net"
	"syscall"
	"time"

	"github.com/miekg/dns"
	"golang.org/x/net/ipv4"

	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
)

func ProbeMDNS(
	iface models.Interface,
) []string {

	logger.Log.Println(
		"MDNS DISCOVERY START",
	)

	networkInterface, err := net.InterfaceByName(
		iface.Name,
	)

	if err != nil {

		logger.Log.Println(
			"MDNS INTERFACE ERROR:",
			err,
		)

		return nil
	}

	logger.Log.Println(
		"MDNS INTERFACE:",
		networkInterface.Name,
		iface.IP,
	)

	target := &net.UDPAddr{
		IP:   net.ParseIP("224.0.0.251"),
		Port: 5353,
	}

	logger.Log.Println(
		"MDNS TARGET:",
		target,
	)

	// mDNS requires UDP port 5353.
	//
	// SO_REUSEADDR allows the socket to coexist with
	// other mDNS listeners such as Avahi.
	//
	// SO_REUSEPORT is also required on Linux when multiple
	// processes need to bind the same UDP port.
	config := net.ListenConfig{
		Control: func(
			network string,
			address string,
			conn syscall.RawConn,
		) error {

			var controlErr error

			err := conn.Control(
				func(fd uintptr) {

					controlErr = syscall.SetsockoptInt(
						int(fd),
						syscall.SOL_SOCKET,
						syscall.SO_REUSEADDR,
						1,
					)

					if controlErr != nil {
						return
					}

					// Linux SO_REUSEPORT = 15.
					controlErr = syscall.SetsockoptInt(
						int(fd),
						syscall.SOL_SOCKET,
						15,
						1,
					)
				},
			)

			if err != nil {
				return err
			}

			return controlErr
		},
	}

	packetConn, err := config.ListenPacket(
		nil,
		"udp4",
		":5353",
	)

	if err != nil {

		logger.Log.Println(
			"MDNS UDP LISTEN ERROR:",
			err,
		)

		return nil
	}

	defer packetConn.Close()

	udpConn := packetConn.(*net.UDPConn)

	err = joinMDNSGroup(
		udpConn,
		networkInterface,
	)

	if err != nil {

		logger.Log.Println(
			"MDNS MULTICAST JOIN ERROR:",
			err,
		)

		return nil
	}

	// PTR query:
	//
	// _services._dns-sd._udp.local
	//
	// This asks mDNS devices to announce their
	// available DNS-SD services.
	request := []byte{
		0x00, 0x00, // Transaction ID
		0x00, 0x00, // Flags
		0x00, 0x01, // Questions
		0x00, 0x00, // Answers
		0x00, 0x00, // Authority
		0x00, 0x00, // Additional

		0x09, '_', 's', 'e', 'r', 'v', 'i', 'c', 'e', 's',
		0x07, '_', 'd', 'n', 's', '-', 's', 'd',
		0x04, '_', 'u', 'd', 'p',
		0x05, 'l', 'o', 'c', 'a', 'l',
		0x00,

		0x00, 0x0c, // PTR
		0x00, 0x01, // IN
	}

	n, err := udpConn.WriteToUDP(
		request,
		target,
	)

	if err != nil {

		logger.Log.Println(
			"MDNS WRITE ERROR:",
			err,
		)

		return nil
	}

	logger.Log.Println(
		"MDNS QUERY SENT:",
		n,
		"bytes",
		target,
	)

	buffer := make(
		[]byte,
		4096,
	)

	deadline := time.Now().Add(
		30 * time.Second,
	)

	found := make(
		map[string]bool,
	)

	for {

		err = udpConn.SetReadDeadline(
			deadline,
		)

		if err != nil {

			logger.Log.Println(
				"MDNS DEADLINE ERROR:",
				err,
			)

			break
		}

		n, addr, err := udpConn.ReadFromUDP(
			buffer,
		)

		if err != nil {

			logger.Log.Println(
				"MDNS READ FINISHED:",
				err,
			)

			break
		}

		msg := new(dns.Msg)

		err = msg.Unpack(
			buffer[:n],
		)

		if err != nil {

			logger.Log.Println(
				"MDNS UNPACK ERROR:",
				err,
			)

			continue
		}

		ip := addr.IP.String()

		// Never add our own host.
		if ip == iface.IP {
			continue
		}

		// Skip duplicate mDNS hosts.
		if found[ip] {
			continue
		}

		found[ip] = true

		logger.Log.Println(
			"MDNS HOST FOUND:",
			ip,
		)
	}

	result := make(
		[]string,
		0,
		len(found),
	)

	for ip := range found {

		result = append(
			result,
			ip,
		)
	}

	logger.Log.Println(
		"MDNS DISCOVERY FOUND:",
		len(result),
	)

	logger.Log.Println(
		"MDNS DISCOVERY FINISHED",
	)

	return result
}

func joinMDNSGroup(
	conn *net.UDPConn,
	networkInterface *net.Interface,
) error {

	packetConn := ipv4.NewPacketConn(
		conn,
	)

	group := &net.UDPAddr{
		IP: net.ParseIP("224.0.0.251"),
	}

	err := packetConn.JoinGroup(
		networkInterface,
		group,
	)

	if err != nil {
		return err
	}

	err = packetConn.SetMulticastInterface(
		networkInterface,
	)

	if err != nil {
		return err
	}

	return nil
}
