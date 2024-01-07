package main

import (
	"GoChiLeMa/src/mainloop"
	"embed"

	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
)

var resources embed.FS

func main() {
	cef.GlobalInit(nil, &resources)
	app := cef.NewApplication()
	app.SetEnableGPU(true)

	config := cef.BrowserWindow.Config
	config.Title = "吃了吗 v1.0.0"
	if common.IsLinux() {
		config.Icon = "resources/img/icon.png"
	} else {
		config.Icon = "resources/img/icon.ico"
	}
	config.Width = 1200
	config.Height = 800
	config.MinWidth = 600
	config.MinHeight = 400
	config.EnableDragFile = true
	config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         &resources,
	}.Build())

	chromiumConfig := config.ChromiumConfig()
	chromiumConfig.SetEnableDevTools(true)
	chromiumConfig.SetEnableMenu(false)
	chromiumConfig.SetEnabledJavascript(true)

	cef.BrowserWindow.SetBrowserInit(mainloop.BrowserInit)

	cef.Run(app)
}
