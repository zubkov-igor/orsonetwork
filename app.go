package main

import (
	"context"
	"OrsoNetwork/internal/logger"
	"OrsoNetwork/internal/models"
	"OrsoNetwork/internal/scanner"
)

type App struct {
	scanner *scanner.Scanner
}

func NewApp() *App {
	return &App{
		scanner: scanner.New(),
	}
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