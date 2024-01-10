package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var lastVal int64 = 0
var previousBeats = []time.Time{}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func logLine(msg any) {
	currentTime := time.Now()
	t := currentTime.Format("15:04:05.0000")
	fmt.Printf("[%s] %s\n", t, msg)
}

func getBeatRate() float64 {
	var first = previousBeats[0]
	var total = len(previousBeats)
	var last = previousBeats[total-1]
	var inBetween = last.Sub(first)
	var perSec = float64(total) / inBetween.Seconds()
	var perMin = perSec * 60
	fmt.Printf("Total Beats: %v, seconds: %v, perSec: %v, perMin: %v\n", total, inBetween.Seconds(), perSec, perMin)
	return perMin
}

func toInt(raw []byte) int64 {
	val, err := strconv.ParseInt((string(raw)), 10, 64)
	if err != nil {
		return lastVal
	}
	lastVal = val
	return val
}

func containsAny(s string, a []string) bool {
	for _, e := range a {
		if strings.Contains(s, e) {
			return true
		}
	}
	return false
}

func getDevices() []string {
	entries, err := os.ReadDir("/dev/")

	if err != nil {
		panic(err)
	}

	likely := []string{}
	for _, e := range entries {
		n := e.Name()
		if containsAny(n, []string{"pf", "tty", "disk", "std", "pty", "random", "zero", "profile", "pf", "ptmx", "aes_0", "apfs", "audit", "auto", "Bluetooth", "console", "afsc", "trace", "fbt", "dt", "events", "log", "lock", "null", "monotonic", "perfmon", "nfsclnt", "Flip5", "nsmb"}) || n == "fd" {
			continue
		}

		fmt.Println(n)
		likely = append(likely, n)
	}

	return likely
}

func pairArduino(device string) io.ReadWriteCloser {
	options := serial.OpenOptions{
		PortName:        fmt.Sprintf("/dev/%s", device),
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	port, err := serial.Open(options)
	if err != nil {
		panic(err)
	}
	return port
}

func heartrate(ctx context.Context, port io.ReadWriteCloser, channel chan float64) {
	var beats = false
	var lastBeat = time.Now()

	for {
		buf := make([]byte, 10)
		port.Read(buf)
		val := toInt(bytes.Trim(buf, "\r\n\x00"))
		runtime.EventsEmit(ctx, "sensordata", val)
		if val != 0 {
			if val > int64(AppData.Settings.Threshold) {
				if beats {
					continue
				}

				beats = true
			} else {
				if !beats {
					continue
				}
				if time.Since(lastBeat) < time.Millisecond*300 {
					beats = false
					continue
				}
				lastBeat = time.Now()
				if len(previousBeats) > 5 {
					previousBeats = append(previousBeats[1:], lastBeat)
				} else {
					previousBeats = append(previousBeats, lastBeat)

				}
				channel <- getBeatRate()
				beats = false
			}
		}
	}
}

func (a *App) GetThreshold() int {
	return AppData.Settings.Threshold
}

func (a *App) UpdateThreshold(i int) {
	AppData.Settings.Threshold = i
	logLine(fmt.Sprintf("Updated threshold to %d", i))
}

func (a *App) ReportResult(rate int) string {
	logLine(fmt.Sprintf("Reported rate! - %v", rate))
	AppData.PreviousReports = append(AppData.PreviousReports, PreviousReport{time.Now().Unix(), rate})
	return "ok"
}

func (a *App) GetSettings() ApplicationSettings {
	logLine("GetSettings()")
	fmt.Println(AppData.Settings)
	return AppData.Settings
}

func (a *App) GetReports() []PreviousReport {
	logLine("GetReports()")
	fmt.Println(AppData.PreviousReports)
	return AppData.PreviousReports
}
