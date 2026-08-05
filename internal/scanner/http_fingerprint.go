package scanner

import (
	"net/http"
	"strings"

	"OrsoNetwork/internal/models"
)

type HTTPFingerprintRule struct {
	Needle      string
	Fingerprint string
}

func AnalyzeHTTPFingerprint(
    info *models.HTTPInfo,
    headers http.Header,
    html string,
) {

    lowerHTML := strings.ToLower(html)

    lowerServer := strings.ToLower(
        headers.Get("Server"),
    )

    for _, rule := range httpFingerprintRules {

        if strings.Contains(
            lowerHTML,
            rule.Needle,
        ) {

            if !containsString(
                info.Fingerprint,
                rule.Fingerprint,
            ) {

                info.Fingerprint = append(
                    info.Fingerprint,
                    rule.Fingerprint,
                )
            }
        }

        if strings.Contains(
            lowerServer,
            rule.Needle,
        ) {

            if !containsString(
                info.Fingerprint,
                rule.Fingerprint,
            ) {

                info.Fingerprint = append(
                    info.Fingerprint,
                    rule.Fingerprint,
                )
            }
        }
    }
}


func containsString(
	values []string,
	value string,
) bool {

	for _, item := range values {

		if item == value {
			return true
		}
	}

	return false
}