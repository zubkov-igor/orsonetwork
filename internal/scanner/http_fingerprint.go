package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)


func AnalyzeHTTPFingerprint(
	info *models.HTTPInfo,
	html string,
) {

	lower := strings.ToLower(html)


	rules := map[string]string{

		"boardtype": "router-ui",
		"softwareversion": "router-firmware",
		"firmwareversion": "router-firmware",

		"signin": "router-login",

		"lighttpd": "embedded-web-server",

		"keenetic": "keenetic-router",
		"tplink": "tp-link-router",
		"zyxel": "zyxel-router",
		"huawei": "huawei-router",
		"zte": "zte-router",

	}


	for word, fingerprint := range rules {

		if strings.Contains(lower, word) {

			info.Fingerprint = append(
				info.Fingerprint,
				fingerprint,
			)

		}
	}

}