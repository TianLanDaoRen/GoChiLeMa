package main

import (
	"GoChiLeMa/src/mainloop"
	"embed"

	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, &resources)
	app := cef.NewApplication()
	app.SetEnableGPU(false)
	app.SetDisableWebSecurity(true)
	app.SetDisableSiteIsolationTrials(true)
	app.SetRemoteDebuggingPort(23457)

	config := cef.BrowserWindow.Config
	config.Title = "吃了吗 v2.24.1.8"
	if common.IsLinux() {
		config.Icon = "resources/img/icon.png"
	} else {
		config.Icon = "resources/img/icon.ico"
	}
	config.Url = "http://localhost:23456/index.html"
	config.Width = 1024
	config.Height = 768
	config.MinWidth = 1024
	config.MinHeight = 1024
	config.EnableDragFile = true
	config.EnableHideCaption = true
	chromiumConfig := config.ChromiumConfig()
	chromiumConfig.SetEnableDevTools(true)
	chromiumConfig.SetEnableMenu(false)
	chromiumConfig.SetEnabledJavascript(true)
	chromiumConfig.SetEnableWindowPopup(true)

	//内置静态资源服务的安全key和value设置
	//通过设置AssetsServerHeaderKeyName和AssetsServerHeaderKeyValue在一定程度上保证资源只能在应用内访问，即使在应用外使用正确的IP和端口号也无法访问到资源
	assetserve.AssetsServerHeaderKeyName = "gochilema-485e983f-8b7a-4e3a-9b0a-9e5f5f5f5f5f"
	assetserve.AssetsServerHeaderKeyValue = "gochilema-485e983f-8b7a-4e3a-9b0a-9e5f5f5f5f5f"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		// fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 23456               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		//LocalAssets 指定本地资源支持热更新 - 适用开发或web端源码可以查看
		// server.LocalAssets = fmt.Sprintf("%s/resources", consts.ExePath)
		//Assets 内置资源不支持热更新 - 适用应用发布
		server.Assets = &resources
		go server.StartHttpServer()
	})

	cef.BrowserWindow.SetBrowserInit(mainloop.BrowserInit)
	cef.Run(app)
}
