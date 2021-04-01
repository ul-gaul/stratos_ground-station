package controller

import (
    "github.com/webview/webview"
)

var view webview.WebView

func Initialize(w webview.WebView) {
    view = w
    // TODO set bindings and other stuff
}

