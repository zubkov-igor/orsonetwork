package scanner

import (
    "io"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "time"

    "OrsoNetwork/internal/logger"
    "OrsoNetwork/internal/models"
)

func ScanHTTP(ip string, port int) models.HTTPInfo {

	info := models.HTTPInfo{
		Port: port,
	}

	info.Scheme = "http"

	url := "http://" + ip + ":" + strconv.Itoa(port)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(url)
if err != nil {
    return info
}

defer resp.Body.Close()

info.Server = resp.Header.Get("Server")
info.StatusCode = resp.StatusCode
info.ContentType = resp.Header.Get("Content-Type")

	body, err := io.ReadAll(
		io.LimitReader(resp.Body, 32768),
	)

	if err != nil {
		return info
	}

html := string(body)
lowerHTML := strings.ToLower(html)

keywords := []string{
    "boardtype",
    "softwareversion",
    "firmwareversion",
    "keenetic",
    "tplink",
    "zyxel",
    "huawei",
    "zte",
}


AnalyzeHTTPFingerprint(
    &info,
    html,
)


for _, word := range keywords {

    if strings.Contains(lowerHTML, word) {

        info.Keywords = append(
            info.Keywords,
            word,
        )
    }
}

scriptPattern := regexp.MustCompile(
    `(?is)<script[^>]*>`,
)

scripts := scriptPattern.FindAllString(
    html,
    -1,
)

srcPattern := regexp.MustCompile(
    `(?i)src\s*=\s*"([^"]+)"`,
)

for _, script := range scripts {
    match := srcPattern.FindStringSubmatch(script)

    if len(match) < 2 {
        continue
    }

    info.Scripts = append(
        info.Scripts,
        match[1],
    )

    logger.Log.Println(
        "HTTP SCRIPT:",
        match[1],
    )
}

logger.Log.Println(
    "HTTP HTML SAMPLE:",
    html[:min(len(html), 500)],
)

start := strings.Index(
    lowerHTML,
    "<title>",
)

end := strings.Index(
    lowerHTML,
    "</title>",
)

logger.Log.Println(
	"TITLE INDEX:",
	start,
	end,
)

logger.Log.Println(
	"HAS TITLE:",
	strings.Contains(
		lowerHTML,
		"title",
	),
)

logger.Log.Println(
    "boardtype:",
    strings.Contains(lowerHTML, "boardtype"),
)

if start != -1 &&
   end != -1 &&
   end > start {

    titleRaw := html[start+7:end]

    logger.Log.Printf(
        "TITLE RAW: %q",
        titleRaw,
    )

    info.Title =
        strings.TrimSpace(titleRaw)
}


logger.Log.Println(
    "HTTP FOUND:",
    ip,
    port,
    info.StatusCode,
    info.ContentType,
    info.Server,
    info.Title,
    info.Keywords,
    info.Fingerprint,
)

return info
}