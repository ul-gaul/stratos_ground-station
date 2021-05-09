package main

import (
    _ "embed"
    log "github.com/sirupsen/logrus"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
    "github.com/ul-gaul/stratos_ground-station/backend/config"
    _ "github.com/ul-gaul/stratos_ground-station/backend/config"
    "github.com/ul-gaul/stratos_ground-station/backend/controller"
    "github.com/ul-gaul/stratos_ground-station/backend/pool"
    "github.com/wailsapp/wails"
)

const Title = "Stratos - Ground Station"

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
    defer pool.Release()
    
    app := wails.CreateApp(&wails.AppConfig{
        Width:            config.General.WindowWidth,
        Height:           config.General.WindowHeight,
        Title:            Title,
        JS:               js,
        CSS:              css,
        Colour:           "",
        Resizable:        true,
        DisableInspector: false,
    })
    
    app.Bind(controller.Controller)
    utils.CheckErr(app.Run())
}

func init() {
    log.SetLevel(config.General.LogLevel)
}
