package acquisition

import (
    log "github.com/sirupsen/logrus"
    "time"
    
    "github.com/ul-gaul/stratos_ground-station/acquisition/packet/collector"
    "github.com/ul-gaul/stratos_ground-station/acquisition/persistence/read"
    "github.com/ul-gaul/stratos_ground-station/acquisition/utils"
)

func Start() {
    log.Info("Application started!")
    start := time.Now()
    
    data, err := read.Csv("./generated_data.csv")
    utils.CheckErr(err)
    collector.Append(data...)
    
    // Keeps the goroutine running (TEMPORARY)
    ticker := time.NewTicker(3 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <- ticker.C:
            log.Infof("Running... (uptime: %s)", time.Since(start).Truncate(time.Millisecond))
        default:
        }
    }
}