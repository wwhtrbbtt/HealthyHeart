package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ApplicationSettings struct {
	SaveFile  string
	Threshold int
}

type PreviousReport struct {
	Timestamp int64
	Heartrate int
}

type ApplicationData struct {
	Settings        ApplicationSettings
	PreviousReports []PreviousReport
}

var AppData = ApplicationData{
	Settings: ApplicationSettings{
		SaveFile:  "/Users/peet/.healthyheart.json",
		Threshold: 550,
	},
	PreviousReports: []PreviousReport{},
}

func (a *App) CreateFile(content string) string {
	logLine("Creating file!")
	fileLocation, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:            "Speicherort",
		DefaultFilename:  "export.csv",
		DefaultDirectory: "/Users/peet/Desktop/",
	})
	logLine("Creating file @ " + fileLocation)

	if err != nil {
		logLine("Error: " + err.Error())
		return err.Error()
	}

	os.WriteFile(fileLocation, []byte(content), 0644)

	return fileLocation
}

func (d ApplicationData) SaveData() {
	logLine("Saving application data")

	j, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		logLine("Error saving application data")
		fmt.Println(err)
	}
	os.WriteFile(d.Settings.SaveFile, j, 0644)
}

func (d ApplicationData) LoadData() {
	logLine("Loading application data")
	dat, err := os.ReadFile(d.Settings.SaveFile)
	if err != nil {
		logLine("Error loading application data!")
		fmt.Println(err)
		d.SaveData()
		return
	}
	var parsed ApplicationData
	json.Unmarshal(dat, &parsed)
	AppData = parsed
}
