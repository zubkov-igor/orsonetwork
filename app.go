package main

import (
    "context"

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

func (a *App) GetTopology() models.Topology {

    topology := a.scanner.Topology()

    println(
        "WAILS:",
        len(topology.Nodes),
        len(topology.Links),
        len(topology.Networks),
    )

    return topology
}
