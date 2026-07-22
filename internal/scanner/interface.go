package scanner

import (
	"net"

	"OrsoNetwork/internal/models"
)

func GetInterfaces() []models.Interface {

	result := []models.Interface{}

	// Получаем список всех сетевых интерфейсов ОС.
	interfaces, err := net.Interfaces()

	if err != nil {
		return result
	}

	for _, iface := range interfaces {

		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if IsVirtualInterface(iface.Name) {
			continue
		}

		addrs, err := iface.Addrs()

		if err != nil {
			continue
		}

		for _, addr := range addrs {

			ipNet, ok := addr.(*net.IPNet)

			if !ok {
				continue
			}

			ip := ipNet.IP

			if ip.To4() == nil {
				continue
			}

			if ip.IsLoopback() {
				continue
			}

			result = append(
				result,
				models.Interface{
					Name:      iface.Name,
					IP:        ip.String(),
					Scannable: true,
				},
			)
		}
	}

	return result
}
