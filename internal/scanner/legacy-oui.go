package scanner

import (
	_ "embed"
	"encoding/json"

	"OrsoNetwork/internal/logger"
	"time"
)

//go:embed data/legacy-oui.json
var legacyJSON []byte

type LegacyVendorRecord struct {
	MACPrefix  string `json:"macPrefix"`
	VendorName string `json:"vendorName"`
}

func loadLegacyVendorDB() {

	start := time.Now()

	logger.Log.Println(
		"LEGACY DB START:",
		len(legacyJSON),
	)

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

	logger.Log.Println(
		"LEGACY DB LOADED:",
		count,
		"TIME:",
		time.Since(start),
	)
}
