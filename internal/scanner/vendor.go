package scanner

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"strings"
	"sync"

	"OrsoNetwork/internal/logger"
)

//go:embed data/macaddress.io-db.json
var vendorJSON []byte

var vendorOnce sync.Once
var legacyOnce sync.Once

var vendors = make(map[string]string)
var legacyVendors = make(map[string]string)

type VendorRecord struct {
	OUI         string `json:"oui"`
	CompanyName string `json:"companyName"`
}

func LookupVendor(mac string) string {

	logger.Log.Println(
		"VENDOR LOOKUP START:",
		mac,
	)

	defer logger.Log.Println(
		"VENDOR LOOKUP END:",
		mac,
	)

	if mac == "" {
		return ""
	}

	vendorOnce.Do(func() {
		loadVendorDB()
	})

	legacyOnce.Do(func() {
		loadLegacyVendorDB()
	})

	prefix := normalizeMAC(mac)

	vendor, ok := vendors[prefix]

	if ok {

		if !strings.Contains(
			vendor,
			"REDACTED",
		) {
			return vendor
		}

	}

	legacyVendor, ok := legacyVendors[prefix]

	if ok {
		return legacyVendor
	}

	return "Unknown"
}

func loadVendorDB() {

	logger.Log.Println(
		"LOADING VENDOR DB BYTES:",
		len(vendorJSON),
	)

	vendors = make(
		map[string]string,
	)

	scanner := bufio.NewScanner(
		bytes.NewReader(vendorJSON),
	)

	count := 0

	for scanner.Scan() {

		var record VendorRecord

		err := json.Unmarshal(
			scanner.Bytes(),
			&record,
		)

		if err != nil {
			logger.Log.Println(
				"VENDOR JSON ERROR:",
				err.Error(),
			)
			continue
		}

		prefix := normalizePrefix(
			record.OUI,
		)

		if prefix == "" {
			continue
		}

		vendors[prefix] = record.CompanyName

		count++
	}

	if err := scanner.Err(); err != nil {

		logger.Log.Println(
			"VENDOR SCANNER ERROR:",
			err.Error(),
		)

		panic(err)
	}

	logger.Log.Println(
		"VENDOR DB LOADED:",
		count,
		"MAP SIZE:",
		len(vendors),
	)
}

func normalizeMAC(mac string) string {

	mac = strings.ToUpper(mac)

	mac = strings.ReplaceAll(mac, ":", "")
	mac = strings.ReplaceAll(mac, "-", "")
	mac = strings.ReplaceAll(mac, ".", "")

	if len(mac) < 6 {
		return ""
	}

	return mac[:6]
}

func normalizePrefix(prefix string) string {

	return normalizeMAC(prefix)
}
