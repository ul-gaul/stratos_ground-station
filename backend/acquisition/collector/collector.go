package collector

import (
    "github.com/olebedev/emitter"
    log "github.com/sirupsen/logrus"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/packet"
)

const (
    EvData  = "data"
)

var Emitter = &emitter.Emitter{}

func Data(packets ...packet.TransformedData) {
    if len(packets) == 0 { return }
    Emitter.Emit(EvData, packets)
}

func OnData(callback func(data []packet.TransformedData)) {
    Emitter.On(EvData, func(ev *emitter.Event) {
        log.Info("DATA EVENT TRIGGERED!")
        callback(ev.Args[0].([]packet.TransformedData))
    })
}