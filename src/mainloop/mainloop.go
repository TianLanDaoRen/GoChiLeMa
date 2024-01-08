package mainloop

import (
	"GoChiLeMa/src/location"
	"GoChiLeMa/src/weather"
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl/version"
)

type LocalStorage struct {
	ip       string
	lat, lon float64
}

var localStorage LocalStorage

func onRequestWeather() {
	go (func() {
		if localStorage.ip == "" {
			var err error
			localStorage.ip, err = location.GetExternalIPv4()
			if err != nil {
				println("Failed to get local IP address:", err)
				return
			}
		}
		ipc.Emit("receiveIp", localStorage.ip)

		go (func() {
			location, err := location.GetLocationInfo(localStorage.ip)
			if err != nil {
				println("Failed to get location information:", err)
				return
			}
			localStorage.lat = location.Lat
			localStorage.lon = location.Lon
			city := fmt.Sprintf("%s %s %s", location.Country, location.RegionName, location.City)
			lat := fmt.Sprintf("%f", localStorage.lat)
			lon := fmt.Sprintf("%f", localStorage.lon)
			timezone := location.Timezone
			isp := location.Isp
			as := location.As
			ipc.Emit("receiveLocation", city, lat, lon, timezone, isp, as)
		})()

		go (func() {
			for localStorage.lat == 0 || localStorage.lon == 0 {
				time.Sleep(time.Millisecond)
			}
			weatherApi := weather.WeatherApi{
				Lat:     fmt.Sprintf("%f", localStorage.lat),
				Lon:     fmt.Sprintf("%f", localStorage.lon),
				Exclude: "minutely,hourly,daily,alerts",
				ApiKey:  "9dc2561b8fe88e3cb9697cfcd4bd770d",
			}
			weatherInfo, ok := weatherApi.GetWeather()
			if !ok {
				println("Failed to get weather information")
				return
			}
			timezoneOffset := int(weatherInfo.TimezoneOffset / 3600)
			temp := fmt.Sprintf("%.2f", weatherInfo.Current.Temp)
			desc := weatherInfo.Current.Weather[0].Description
			humidity := int(weatherInfo.Current.Humidity)
			uvi := weatherInfo.Current.Uvi
			clouds := int(weatherInfo.Current.Clouds)
			windSpeed := weatherInfo.Current.WindSpeed
			sunriseTimestamp := int64(weatherInfo.Current.Sunrise)
			sunrise := time.Unix(sunriseTimestamp, 0).UTC().Add(time.Hour * time.Duration(timezoneOffset)).Format("15:04:05")
			sunsetTimestamp := int64(weatherInfo.Current.Sunset)
			sunset := time.Unix(sunsetTimestamp, 0).UTC().Add(time.Hour * time.Duration(timezoneOffset)).Format("15:04:05")
			description := fmt.Sprintf("%s，湿度%d%%，紫外线指数%.2f，云度%d%%，风速%.2fm/s，日出时间%s，日落时间%s", desc, humidity, uvi, clouds, windSpeed, sunrise, sunset)
			ipc.Emit("receiveWeather", temp, description)
		})()
	})()
}

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

	ipc.On("requestContactUs", func() {
		ipc.Emit("receiveAllowContactUs", "moriartylimitter@outlook.com", "TianLanDaoRen")
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

	ipc.On("requestWeather", onRequestWeather)
}

func MainLoop() {
	for {
		ipc.Emit("receiveTick")
		time.Sleep(time.Second)
	}
}

func BrowserInit(event *cef.BrowserEvent, bw cef.IBrowserWindow) {
	listenOnIpc()
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
