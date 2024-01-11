package main

import (
	"GoChiLeMa/src/config"
	"GoChiLeMa/src/mainloop"
	"embed"

	"github.com/energye/energy/v2/cef"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, &resources)
	app := cef.NewApplication()
	app.SetEnableGPU(false)
	app.SetDisableWebSecurity(true)
	if config.DEBUGING {
		app.SetRemoteDebuggingPort(config.DEBUGING_PORT)
	}

	mainloop.SetupBrowserWindowConfig()
	mainloop.SetupHttpService(&resources)

	cef.BrowserWindow.SetBrowserInit(mainloop.BrowserInit)
	cef.Run(app)
}
