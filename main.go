package main

import (
    _ "embed"
    log "github.com/sirupsen/logrus"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
    "github.com/ul-gaul/stratos_ground-station/backend/controller"
    "github.com/ul-gaul/stratos_ground-station/backend/pool"
    "github.com/wailsapp/wails"
)

// TODO obtain from config
const (
    Title         = "Stratos - Base Station"
    DefaultWidth  = 1024
    DefaultHeight = 768
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
    defer pool.Release()
    
    app := wails.CreateApp(&wails.AppConfig{
        Width:            DefaultWidth,
        Height:           DefaultHeight,
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
    log.SetLevel(log.DebugLevel)
}
