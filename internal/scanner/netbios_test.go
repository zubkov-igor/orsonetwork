package scanner

import (
	"log"
	"testing"
)

func TestNetBIOSLookup(t *testing.T) {

	result, err := LookupNetBIOS(
		"192.168.0.37",
	)

	if err != nil {
		t.Fatal(err)
	}

	log.Println(
		"NETBIOS RESULT:",
		result,
	)
}