// +build !dev

package main

import (
    "embed"
    "fmt"
    log "github.com/sirupsen/logrus"
    "github.com/webview/webview"
    "net"
    "net/http"
    
    "github.com/ul-gaul/stratos_ground-station/acquisition"
    "github.com/ul-gaul/stratos_ground-station/acquisition/utils"
    "github.com/ul-gaul/stratos_ground-station/controller"
    "github.com/ul-gaul/stratos_ground-station/pool"
)

// Set via -ldflags
var FrontendIndexPath string

//go:embed dashboard/dist
var content embed.FS

func main() {
    defer pool.Release()
    done := make(chan struct{})
    
    log.Infof("FrontendIndexPath: %s", FrontendIndexPath)
    
    utils.CheckErr(pool.Frontend.Submit(func() {
        w := webview.New(true)
        defer w.Destroy()
        
        controller.Initialize(w)
        
        ln, err := net.Listen("tcp", "127.0.0.1:0")
        utils.CheckErr(err)
        defer ln.Close()
        go http.Serve(ln, http.FileServer(http.FS(content)))

        w.Navigate(fmt.Sprintf("http://%s/%s", ln.Addr(), FrontendIndexPath))
        
        done <- struct{}{} // Ready to execute backend
        w.Run()
        done <- struct{}{} // App exited
    }))
    
    <- done // Ready to execute backend
    utils.CheckErr(pool.Backend.Submit(acquisition.Start))
    <- done // App exited
}

