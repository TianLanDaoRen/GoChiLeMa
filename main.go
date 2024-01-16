package main

import (
	"GoChiLeMaWails/src/encrypto"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	crypto := encrypto.NewEncrypto()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "吃了吗 - GoChiLeMa",
		Width:     1024,
		Height:    576,
		MinWidth:  1024,
		MinHeight: 576,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnDomReady:       app.domReady,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 80},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			crypto,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
