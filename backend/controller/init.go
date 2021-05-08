package controller

import (
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/collector"
    . "github.com/ul-gaul/stratos_ground-station/backend/acquisition/packet"
    "github.com/ul-gaul/stratos_ground-station/backend/pool"
    "github.com/wailsapp/wails"
)

var Controller = &controller{}

type controller struct {
    r      *wails.Runtime
    store  *wails.Store
}

// WailsInit is called by wails at initialization
func (c *controller) WailsInit(r *wails.Runtime) error {
    c.r = r
    
    collector.OnData(func(data []TransformedData) {
        r.Events.Emit("data", data)
    })
    
    return pool.Backend.Submit(acquisition.Start)
}

// WailsShutdown is called by wails at shutdown
func (c *controller) WailsShutdown() {
}
