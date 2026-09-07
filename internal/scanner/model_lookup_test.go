package scanner

import (
	"testing"
)

func TestLookupDeviceModel(t *testing.T) {
	model, ok := LookupDeviceModel("XQ-DE72")

	if !ok {
		t.Fatal("device model not found")
	}

	if model.Brand != "Sony" {
		t.Fatalf("unexpected brand: %s", model.Brand)
	}

	if model.Name != "Xperia 5 V" {
		t.Fatalf("unexpected model name: %s", model.Name)
	}
}


func TestLookupDeviceModelNotFound(t *testing.T) {
	_, ok := LookupDeviceModel("THIS-IS-NOT-A-DEVICE-CODE")

	if ok {
		t.Fatal("unexpected device model match")
	}
}



func TestFindDeviceModelFromHostname(t *testing.T) {
      model, ok := FindDeviceModelFromHostname(
            "xq-de72.igd_rostelecom",
      )

      if !ok {
            t.Fatal("device model not found")
      }

      if model.Brand != "Sony" {
            t.Fatalf("unexpected brand: %s", model.Brand)
      }

      if model.Name != "Xperia 5 V" {
            t.Fatalf("unexpected model: %s", model.Name)
      }
}


func TestFindDeviceModelFromHostnameNoGuess(t *testing.T) {
	hostnames := []string{
		"a55-pol-zovatela-zubkova.igd_rostelecom",
		"m2004j19c-redmi9.igd_rostelecom",
		"atom32.igd_rostelecom",
	}

	for _, hostname := range hostnames {
		_, ok := FindDeviceModelFromHostname(hostname)

		if ok {
			t.Fatalf("unexpected device model match for hostname: %s", hostname)
		}
	}
}