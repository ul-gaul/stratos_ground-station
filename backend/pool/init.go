package pool

import (
    "github.com/panjf2000/ants/v2"
    log "github.com/sirupsen/logrus"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
)

var Frontend *ants.Pool
var Backend *ants.Pool

func init() {
    var err error
    
    Frontend, err = ants.NewPool(1, ants.WithLogger(log.New()))
    utils.CheckErr(err)
    
    Backend, err = ants.NewPool(ants.DefaultAntsPoolSize, ants.WithLogger(log.New()))
    utils.CheckErr(err)
}

func Release() {
    defer Backend.Release()
    defer Frontend.Release()
    defer ants.Release()
}