package mainloop

import (
	CONFIG "GoChiLeMa/src/config"
	"GoChiLeMa/src/route"
	"embed"
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl/version"
)

func listenOnIpc() {
	ipc.On("requestWindowState", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		state := context.ArgumentList().GetIntByIndex(0)
		if state == 0 {
			bw.Minimize()
		} else if state == 1 {
			bw.Maximize()
		} else if state == 3 {
			if bw.IsFullScreen() {
				bw.ExitFullScreen()
			} else {
				bw.FullScreen()
			}
		}
	})

	ipc.On("requestWindowClose", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.CloseBrowserWindow()
	})

	ipc.On("requestCloseApp", func() {
		cef.BrowserWindow.MainWindow().CloseBrowserWindow()
		cef.BrowserWindow.MainWindow().Close()
	})

	ipc.On("requestOpenLinkInExternalBrowser", func(url string) {
		cmd := exec.Command("cmd", "/c", "start", url)
		if runtime.GOOS == "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		}
		cmd.Run()
	})
}

func MainLoop() {
	for {
		ipc.Emit("receiveTick")
		time.Sleep(time.Second)
	}
}

func BrowserInit(event *cef.BrowserEvent, bw cef.IBrowserWindow) {
	route.Init()
	listenOnIpc()
	// Set event handlers
	event.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32, window cef.IBrowserWindow) {
		// Set user agent
		info := cef.BrowserWindow.GetWindowInfo(browser.BrowserId())
		dict := &cef.ICefDictionaryValue{}
		dict.SetString("userAgent", "Mozilla/5.0 (Linux; Android 11; M2102K1G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Mobile Safari/537.36")
		info.Chromium().ExecuteDevToolsMethod(0, "", dict)
		// IPC
		ipc.Emit("receiveOsInfo", version.OSVersion.ToString())
		// Set window style
		if window.IsLCL() {
			// 边框圆角
			window.AsLCLBrowserWindow().SetRoundRectRgn(10)
			window.AsLCLBrowserWindow().BrowserWindow().SetRoundRectRgn(10)
			// 更新窗口以显示圆角
			rect := window.Bounds()
			window.SetBounds(rect.X, rect.Y, rect.Width, rect.Height+1)
			window.SetBounds(rect.X, rect.Y, rect.Width, rect.Height)
		}
		go MainLoop()
	})
}

func SetupHttpService(resources *embed.FS) {
	//内置静态资源服务的安全key和value设置
	//通过设置AssetsServerHeaderKeyName和AssetsServerHeaderKeyValue在一定程度上保证资源只能在应用内访问，即使在应用外使用正确的IP和端口号也无法访问到资源
	assetserve.AssetsServerHeaderKeyName = "gochilema-485e983f-8b7a-4e3a-9b0a-9e5f5f5f5f5f"
	assetserve.AssetsServerHeaderKeyValue = "gochilema-485e983f-8b7a-4e3a-9b0a-9e5f5f5f5f5f"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		// fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = CONFIG.PORT         //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		//LocalAssets 指定本地资源支持热更新 - 适用开发或web端源码可以查看
		// server.LocalAssets = fmt.Sprintf("%s/resources", consts.ExePath)
		//Assets 内置资源不支持热更新 - 适用应用发布
		server.Assets = resources
		go server.StartHttpServer()
	})
}

func SetupBrowserWindowConfig() {
	config := cef.BrowserWindow.Config
	config.Title = "吃了吗 v3.24.1.11"
	if common.IsLinux() {
		config.Icon = "resources/img/icon.png"
	} else {
		config.Icon = "resources/img/icon.ico"
	}
	config.Url = fmt.Sprintf("http://localhost:%d/index.html", CONFIG.PORT)
	config.Width = 1024
	config.Height = 768
	config.MinWidth = 1024
	config.MinHeight = 1024
	config.EnableDragFile = true
	config.EnableHideCaption = true
	chromiumConfig := config.ChromiumConfig()
	chromiumConfig.SetEnableDevTools(CONFIG.DEBUGING)
	chromiumConfig.SetEnableMenu(false)
	chromiumConfig.SetEnabledJavascript(true)
	chromiumConfig.SetEnableWindowPopup(true)
}
