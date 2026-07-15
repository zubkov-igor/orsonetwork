package main

import (
	"fmt"
	"time"

	"OrsoNetwork/internal/scanner"
)

func main() {

	fmt.Println("========== OrsoNetwork ==========")

	s := scanner.New()

	start := time.Now()

	networks := s.Scan()

	fmt.Println(
		"Scan time:",
		time.Since(start),
	)

	for _, network := range networks {

		fmt.Println()

		fmt.Println(
			"Interface:",
			network.Interface,
		)

		fmt.Println(
			"Network:",
			network.CIDR,
		)

		fmt.Println(
			"Gateway:",
			network.Gateway,
		)

		fmt.Println(
			"Hosts:",
			len(network.Hosts),
		)

		for _, host := range network.Hosts {

			fmt.Printf(
				"  %-15s %-5v %v\n",
				host.IP,
				host.Online,
				host.RTT,
			)
		}
	}
}