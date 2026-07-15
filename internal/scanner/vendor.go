package scanner

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed data/mac-vendors.json
var vendorJSON []byte

var vendors map[string]string


type VendorRecord struct {
	MACPrefix  string `json:"macPrefix"`
	VendorName string `json:"vendorName"`
	Private    bool   `json:"private"`
	BlockType  string `json:"blockType"`
	LastUpdate string `json:"lastUpdate"`
}


func LookupVendor(mac string) string {

	if mac == "" {
		return ""
	}


	if vendors == nil {
		LoadVendorDB()
	}


	prefix := normalizeMAC(mac)


	vendor, ok := vendors[prefix]


	if ok {
		return vendor
	}


	return "Unknown"
}


func LoadVendorDB() {

	var records []VendorRecord

	err := json.Unmarshal(
		vendorJSON,
		&records,
	)

	if err != nil {
		panic(err)
	}


	vendors = make(
		map[string]string,
	)


	for _, r := range records {

		vendors[r.MACPrefix] = r.VendorName

	}
}


func normalizeMAC(mac string) string {

	mac = strings.ToUpper(mac)

	mac = strings.ReplaceAll(mac, "-", ":")
	mac = strings.ReplaceAll(mac, ".", "")


	if len(mac) < 8 {
		return ""
	}


	return mac[:8]
}