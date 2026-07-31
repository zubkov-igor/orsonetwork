package scanner

import (
    "bufio"
    "bytes"
    _ "embed"
    "strings"
    "sync"

    "OrsoNetwork/internal/logger"
)

//go:embed data/ieee-oui.txt
var vendorTXT []byte

var (
    vendorOnce sync.Once
    vendors    = make(map[string]string)
    dbLoaded   bool // чтобы можно было явно проверить состояние, если нужно
)

func LookupVendor(mac string) string {
    if mac == "" {
        return ""
    }

    vendorOnce.Do(loadVendorDB)

    prefix := normalizeMAC(mac)
    if prefix == "" {
        return "Unknown"
    }

    if vendor, ok := vendors[prefix]; ok {
        return vendor
    }
    return "Unknown"
}

func loadVendorDB() {
    vendors = make(map[string]string)

    logger.Log.Printf("Vendor file bytes: %d", len(vendorTXT))

    if len(vendorTXT) == 0 {
        logger.Log.Println("WARNING: vendorTXT is empty")
        dbLoaded = true
        return
    }

    scanner := bufio.NewScanner(bytes.NewReader(vendorTXT))
    count := 0
    skipped := 0

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        // Пропускаем пустые строки и комментарии
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        fields := strings.Fields(line)
        if len(fields) < 2 {
            skipped++
            continue
        }

        rawPrefix := fields[0]
        prefix := normalizePrefix(rawPrefix)

        if prefix == "" {
            skipped++
            continue
        }

        // В manuf: [префикс] [короткое имя] [полное имя...]
        // Мы хотим полное имя (начиная со 2-го поля).
        // Но если вдруг короткое имя совпадает с полным — тоже ок.
        vendor := strings.Join(fields[2:], " ") // берём полное имя
        if vendor == "" {
            vendor = fields[1] // запасной вариант: короткое имя
        }

        vendors[prefix] = vendor
        count++

       // logger.Log.Printf("[OUI-DEBUG] RAW: %q -> KEY: %s | VENDOR: %s", rawPrefix, prefix, vendor)
    }

    if err := scanner.Err(); err != nil {
        logger.Log.Printf("Scanner error: %v", err)
    }

    logger.Log.Printf("IEEE OUI LOADED: %d entries, skipped %d lines, map size: %d", count, skipped, len(vendors))
    dbLoaded = true
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
    p := normalizeMAC(prefix)
    if len(p) >= 6 {
        return p[:6]
    }
    return ""
}
