package scanner

import (
	"strings"

	"OrsoNetwork/internal/models"
)

var httpKeywords = []string{
	"boardtype",
	"softwareversion",
	"firmwareversion",
	"keenetic",
	"tplink",
	"zyxel",
	"huawei",
	"zte",
}

func AnalyzeHTTPKeywords(
	info *models.HTTPInfo,
	html string,
) {

	lower := strings.ToLower(html)

	for _, keyword := range httpKeywords {

		if strings.Contains(lower, keyword) {

			info.Keywords = append(
				info.Keywords,
				keyword,
			)
		}
	}
}
