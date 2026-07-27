package scanner

func PortService(port int) string {

	switch port {

	case 22:
		return "ssh"

	case 53:
		return "dns"

	case 80:
		return "http"

	case 139:
		return "netbios"

	case 443:
		return "https"

	case 445:
		return "smb"

	case 515:
		return "lpd"

	case 554:
		return "rtsp"

	case 631:
		return "ipp"

	case 9100:
		return "jetdirect"

	case 3389:
		return "rdp"

	default:
		return "unknown"
	}
}