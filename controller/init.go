package controller

import (
    "encoding/json"
    "fmt"
    emit "github.com/olebedev/emitter"
    log "github.com/sirupsen/logrus"
    "github.com/webview/webview"
    
    . "github.com/ul-gaul/stratos_ground-station/acquisition/packet"
    "github.com/ul-gaul/stratos_ground-station/acquisition/packet/collector"
    "github.com/ul-gaul/stratos_ground-station/acquisition/utils"
)

// TODO obtain from config
const (
    Title        = "Stratos - Base Station"
    DefaultWidth  = 800
    DefaultHeight = 800
)

const (
    EvConnect = "connect"
)

var emitter *emit.Emitter
var view webview.WebView

func init() {
    emitter = &emit.Emitter{}
}

func Emitter() *emit.Emitter {
    return emitter
}

func Initialize(w webview.WebView) {
    if view != nil {
        log.Warn("controller initialization ignored: already initialized")
        return
    }
    
    if w == nil {
        log.Fatal("controller initialized with nil")
    }
    
    view = w
    view.SetTitle(Title)
    view.SetSize(DefaultWidth, DefaultHeight, webview.HintNone)
    view.Init(`window.go_ondata = window.go_ondata || (function() {})`)
    view.Eval(`window.go_ondata = window.go_ondata || (function() {})`)
    
    utils.CheckErr(view.Bind("go_connect", onConnect))
    
    collector.Collector().On(collector.EvAppend, func(ev *emit.Event) {
        utils.CheckErr(toFrontend(ev.Args[0].([]TransformedData)))
    })
}

// toFrontend sends the data to the frontend
func toFrontend(packets []TransformedData) error {
    jsonData, err := json.Marshal(packets)
    if err != nil { return err }
    
    view.Dispatch(func() {
        view.Eval(fmt.Sprintf("window.go_ondata(%s)", jsonData))
    })
    
    return nil
}


func onConnect() ([]TransformedData, error) {
    emitter.Emit(EvConnect)
    return collector.Data(), nil
}