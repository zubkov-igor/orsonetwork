package models

type DeviceType string

const (
	DeviceUnknown DeviceType = "unknown"

	DeviceRouter DeviceType = "router"
	DeviceGateway DeviceType = "gateway"

	DeviceComputer DeviceType = "computer"
	DeviceServer DeviceType = "server"

	DevicePhone DeviceType = "phone"
	DeviceTablet DeviceType = "tablet"

	DevicePrinter DeviceType = "printer"

	DeviceCamera DeviceType = "camera"

	DeviceNAS DeviceType = "nas"

	DeviceTV DeviceType = "tv"
	DeviceConsole DeviceType = "console"

	DeviceAccessPoint DeviceType = "access_point"

	DeviceIoT DeviceType = "iot"
)