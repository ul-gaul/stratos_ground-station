package collector

import (
    "github.com/olebedev/emitter"
    "sync"
    
    . "github.com/ul-gaul/stratos_ground-station/acquisition/packet"
)

var c *collector

const (
    EvAppend = "append"
    EvClear = "clear"
)

func init() {
    c = &collector{}
    c.Emitter = &emitter.Emitter{}
}

type collector struct {
    *emitter.Emitter
    sync.Mutex
    data []TransformedData
}

func Collector() *collector {
    return c
}

// Append adds the packets at the end of the slice
func Append(packets ...TransformedData) {
    if len(packets) == 0 { return }
    c.Lock()
    c.data = append(c.data, packets...)
    c.Unlock()
    c.Emit(EvAppend, packets)
}

// Clear resets the slice of data
func Clear() {
    c.Lock()
    c.data = nil
    c.Unlock()
    c.Emit(EvClear)
}

// Data returns the slice of data
func Data() []TransformedData {
    return c.data
}
