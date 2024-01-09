package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var connectedDevice = false

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	AppData.LoadData()
}

func (a *App) domReady(ctx context.Context) {
	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			if connectedDevice {
				continue
			}
			connect(ctx)
		}
	}()
}

func (a *App) IsConnected() bool {
	return connectedDevice
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	AppData.SaveData()
	return false
}

func (a *App) shutdown(ctx context.Context) {}

func connect(ctx context.Context) {
	devices := getDevices()
	logLine("Found: " + strings.Join(devices, ", "))
	if len(devices) != 1 {
		connectedDevice = false
		return
	}
	connectedDevice = true
	fmt.Println("Pairing with arduino", devices[0])
	monitor := pairArduino(devices[0])
	rate := make(chan float64)
	go heartrate(ctx, monitor, rate)
	go func() {
		for {
			if !connectedDevice {
				break
			}
			curr := <-rate
			runtime.EventsEmit(ctx, "heartrate", curr)
			fmt.Println(curr)
		}
	}()
}
