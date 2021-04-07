package packet

import (
    "time"
)

// TransformedData represent the human readable form of data
type TransformedData struct {
    Time time.Duration `json:"time" csv:"time"`
    Latitude float64 `json:"latitude" csv:"latitude"`
    Longitude float64 `json:"longitude" csv:"longitude"`
    Altitude float64 `json:"altitude" csv:"altitude"`
    Pressure float64 `json:"pressure" csv:"pressure"`
    InternalTemp float64 `json:"internalTemp" csv:"internal_temp"`
    ExternalTemp float64 `json:"externalTemp" csv:"external_temp"`
}