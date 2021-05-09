package config

import (
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "go.bug.st/serial"
)

func applyDefaults() {
    // General
    viper.SetDefault("General.LogLevel", log.InfoLevel)
    viper.SetDefault("General.WindowWidth", 1024)
    viper.SetDefault("General.WindowHeight", 768)
    
    // Frontend
    viper.SetDefault("Frontend.DefaultDataLength", 50)
    
    // Comms
    viper.SetDefault("Comms.UseBigEndian", false)
    
    viper.SetDefault("Comms.Serial.BaudRate", 115_200)
    viper.SetDefault("Comms.Serial.DataBits", 8)
    viper.SetDefault("Comms.Serial.Parity", serial.NoParity)
    viper.SetDefault("Comms.Serial.StopBits", serial.OneStopBit)
    
    viper.SetDefault("Comms.DataPacket.LossThreshold", 20)
    viper.SetDefault("Comms.DataPacket.BufferSize", 5)
}

var configs = map[string]interface{}{
    "General":  &General,
    "Frontend": &Frontend,
    "Comms":    &Comms,
}


// General settings
var General struct {
    // LogLevel is the verbosity (default: log.InfoLevel)
    LogLevel log.Level
    
    // WindowWidth is the default width of the application window (default: 1024)
    WindowWidth int
    
    // WindowWidth is the default height of the application window (default: 768)
    WindowHeight int
}

// Frontend contains the frontend settings
var Frontend struct {
    
    // DefaultDataLength is the default maximum length of data shown in the
    // charts and in the table. (default: 50)
    DefaultDataLength uint
}

// Comms contains the settings related to communications.
var Comms struct {
    
    // UseBigEndian specifies the endianness (or byte order) of the packets.
    // (default: false)
    UseBigEndian bool
    
    // DataPacket contains the settings for the reception of data packets.
    DataPacket struct {
        // LossThreshold is the limit of consecutive corrupted data packets to
        // ignore before it crashes. (default: 20)
        LossThreshold uint
        
        // BufferSize is the size of the channel receiving the data packets.
        // (default: 5)
        BufferSize uint
    }
    
    // Serial contains the configuration of the serial port
    Serial serial.Mode
}
