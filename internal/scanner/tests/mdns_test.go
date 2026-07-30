//go:build mdns

package scanner

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/grandcat/zeroconf"
)

func TestMDNS(t *testing.T) {

	services := []string{
		"_device-info._tcp",
		"_airplay._tcp",
		"_googlecast._tcp",
		"_ssh._tcp",
		"_smb._tcp",
	}

	for _, service := range services {

		log.Println(
			"SEARCH MDNS:",
			service,
		)

		resolver, err := zeroconf.NewResolver(nil)

		if err != nil {
			t.Fatal(err)
		}

		entries := make(chan *zeroconf.ServiceEntry)

		go func() {
			for entry := range entries {

				log.Println(
					"MDNS ENTRY:",
					entry.Instance,
					entry.HostName,
					entry.AddrIPv4,
				)
			}
		}()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			3*time.Second,
		)

		err = resolver.Browse(
			ctx,
			service,
			"local.",
			entries,
		)

		if err != nil {
			t.Log(
				"MDNS ERROR:",
				service,
				err,
			)
		}

		<-ctx.Done()

		cancel()
	}
}

func TestMDNSLookup(t *testing.T) {

	resolver, err := zeroconf.NewResolver(nil)

	if err != nil {
		t.Fatal(err)
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func() {

		for entry := range entries {

			log.Println(
				"LOOKUP ENTRY:",
				entry.Instance,
				entry.HostName,
				entry.AddrIPv4,
			)
		}

	}()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	err = resolver.Lookup(
		ctx,
		"iPhone.local",
		entries,
	)

	if err != nil {
		t.Fatal(err)
	}

	<-ctx.Done()
}
