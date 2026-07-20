package scanner

import "strings"

func IsVirtualInterface(name string) bool {

	name = strings.ToLower(name)

	blacklist := []string{
		"radmin",
		"virtual",
		"hyper",
		"vmware",
		"vmnet",
		"virtualbox",
		"teredo",
		"tap",
		"tun",
		"docker",
		"veth",
		"virbr",
	}

	for _, b := range blacklist {

		if strings.Contains(name, b) {
			return true
		}
	}

	return false
}
