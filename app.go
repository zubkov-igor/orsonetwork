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
