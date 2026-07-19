package scanner

import (
	"encoding/json"
	_ "embed"
)


//go:embed data/legacy-oui.json
var legacyJSON []byte


type LegacyVendorRecord struct {
	MACPrefix  string `json:"macPrefix"`
	VendorName string `json:"vendorName"`
}


func LoadLegacyVendorDB() {

    var records []LegacyVendorRecord


    err := json.Unmarshal(
        legacyJSON,
        &records,
    )


    if err != nil {
        panic(err)
    }


    legacyVendors = make(
        map[string]string,
    )


    count := 0


    for _, record := range records {


        prefix := normalizePrefix(
            record.MACPrefix,
        )


        if prefix == "" {
            continue
        }


        legacyVendors[prefix] =
            record.VendorName


        count++
    }


    println(
        "LEGACY DB LOADED:",
        count,
    )
}