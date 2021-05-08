package acquisition

import (
    log "github.com/sirupsen/logrus"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/collector"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/packet"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/persistence/read"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
    "time"
)

func Start() {
    log.Debug("Backend running...")
    for {
        select {
        case <-time.After(5 * time.Second):
            log.Debug("Backend running...")
        default:
        }
    }
}

func TestData() []packet.TransformedData {
    log.Info("Reading data from file...")
    data, err := read.Csv("./generated_data.csv")
    utils.CheckErr(err)
    return data
}

func SendTestData() {
    log.Info("Sending data...")
    collector.Data(TestData()...)
    log.Info("Data sent!")
}