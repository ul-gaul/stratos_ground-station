package acquisition

import (
    log "github.com/sirupsen/logrus"
    "time"
)

func Start() {
    // TODO Récupérer les données
    log.Info("Application started!")
    start := time.Now()
    
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