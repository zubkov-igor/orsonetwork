package scanner

import "testing"

func TestLookupVendorFormats(t *testing.T) {

	tests := []string{
		"ec:b1:e0:0f:23:90",
		"EC:B1:E0:0F:23:90",
		"ec-b1-e0-0f-23-90",
	}

	for _, mac := range tests {

		vendor := LookupVendor(mac)

		t.Log(
			mac,
			"->",
			vendor,
		)

	}

}
