package main

import (
	"GoChiLeMaWails/src/route"
	"context"
	"os/exec"
	"runtime"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Init route
	route.Init()
}

// domReady is called when the frontend is ready to receive messages
func (a *App) domReady(ctx context.Context) {

}

// Open link in external browser
func (a *App) OpenLinkInExternalBrowser(url string) {
	cmd := exec.Command("cmd", "/c", "start", url)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	cmd.Run()
}
