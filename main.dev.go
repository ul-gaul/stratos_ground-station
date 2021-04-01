// +build dev

package main

import (
    "fmt"
    "github.com/webview/webview"
    "os"
    
    "github.com/ul-gaul/stratos_ground-station/acquisition"
    "github.com/ul-gaul/stratos_ground-station/acquisition/utils"
    "github.com/ul-gaul/stratos_ground-station/controller"
    "github.com/ul-gaul/stratos_ground-station/pool"
)

const DefaultPort = "8080"

func main() {
    defer pool.Release()
    done := make(chan struct{})
    
    utils.CheckErr(pool.Frontend.Submit(func() {
        w := webview.New(true)
        defer w.Destroy()
    
        controller.Initialize(w)
    
        port := os.Getenv("PORT")
        if port == "" {
            port = DefaultPort
        }
        w.Navigate(fmt.Sprintf("http://127.0.0.1:%s/", port))
    
        done <- struct{}{} // Ready to execute backend
        w.Run()
        done <- struct{}{} // App exited
    }))
    
    <- done // Ready to execute backend
    utils.CheckErr(pool.Backend.Submit(acquisition.Start))
    <- done // App exited
}