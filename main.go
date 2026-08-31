package main

import (
	"embed"

	"OrsoNetwork/internal/logger"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

    logger.Init()

    // Log levels:
    //
    // LevelDebug — подробная информация для разработки и отладки.
    // LevelInfo  — основные события и результаты работы программы.
    // LevelWarn  — потенциальные проблемы, которые не остановили работу.
    // LevelError — ошибки, из-за которых операция не выполнена.
    //
    // Current level:

    logger.SetLevel(
        logger.LevelInfo,
    )


	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "OrsoNetwork",

		WindowStartState: options.Maximised,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		BackgroundColour: &options.RGBA{
			R: 27,
			G: 38,
			B: 54,
			A: 1,
		},

		OnStartup: app.startup,

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
