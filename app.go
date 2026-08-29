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

	logger.Log.Println(
		"TOPOLOGY:",
		len(topology.Topology.Nodes),
		len(topology.Topology.Links),
		len(topology.Topology.Networks),
	)

	return topology
}

func (a *App) UpdateScannerConfig(
    config models.ScannerConfig,
) {

    logger.Log.Println(
        "NEW SCANNER CONFIG:",
        config,
    )

    a.scanner.UpdateConfig(
        config,
    )
}