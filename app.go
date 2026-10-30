package main

import (
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
	"OrsoNetwork/internal/scanner"
	"OrsoNetwork/internal/wifi"
	"context"
	"time"
)

type App struct {
	scanner *scanner.Scanner
}

type DeepScanResult struct {
	AccessPoints []wifi.AccessPoint `json:"accessPoints"`
	Clients      []wifi.Client      `json:"clients"`
}

func NewApp() *App {
	return &App{
		scanner: scanner.New(),
	}
}

func (a *App) DeepScanWifi() (DeepScanResult, error) {
	iface, err := wifi.FindWirelessInterface()
	if err != nil {
		return DeepScanResult{}, err
	}

	accessPoints, clients, err := wifi.DeepScan(
		context.Background(),
		iface,
		10*time.Second,
	)
	if err != nil {
		return DeepScanResult{}, err
	}

	for i := range clients {
		host, ok := a.scanner.FindHostByMAC(clients[i].MAC)
		if !ok {
			continue
		}

		logger.Info(
			"WIFI CLIENT MATCH:",
			clients[i].MAC,
			"IP:",
			host.IP,
			"HOSTNAME:",
			host.Hostname,
		)

		model, modelOK := scanner.FindDeviceModelFromHostname(host.Hostname)

		if modelOK {
			clients[i].Brand = model.Brand
			clients[i].Model = model.Name

			logger.Info(
				"WIFI CLIENT MODEL:",
				clients[i].MAC,
				"BRAND:",
				model.Brand,
				"MODEL:",
				model.Name,
			)
		}
	}

	return DeepScanResult{
		AccessPoints: accessPoints,
		Clients:      clients,
	}, nil
}

func (a *App) startup(ctx context.Context) {

}

func (a *App) GetTopology() models.ScanResult {

	topology := a.scanner.Topology()

	logger.Info(
		"TOPOLOGY:",
		len(topology.Topology.Nodes),
		len(topology.Topology.Links),
		len(topology.Topology.Networks),
	)

	return topology
}

func (a *App) UpdateScannerConfig(
	config models.ScannerConfig,
) error {

	logger.Info(
		"NEW SCANNER CONFIG:",
		config,
	)

	return a.scanner.UpdateConfig(
		config,
	)
}

func (a *App) GetScannerConfig() models.ScannerConfig {
	return a.scanner.GetConfig()
}

func (a *App) ScanWifi() ([]wifi.Network, error) {
	return wifi.Scan()
}
