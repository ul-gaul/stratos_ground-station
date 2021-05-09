package config

import (
    "github.com/mitchellh/go-homedir"
    "github.com/mitchellh/mapstructure"
    "github.com/spf13/viper"
    "github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
)

const (
    envPrefix = "STRATOS"
    cfgName   = "stratos-config"
)

func init() {
    var err error
    var path string
    
    // Find home directory
    path, err = homedir.Dir()
    utils.CheckErr(err)
    viper.AddConfigPath(path)
    
    // Find executable path
    path, err = utils.GetExecutableDir()
    utils.CheckErr(err)
    viper.AddConfigPath(path)
    
    viper.AddConfigPath(".")
    viper.SetConfigName(cfgName)
    viper.SetEnvPrefix(envPrefix)
    viper.AutomaticEnv()
    
    applyDefaults()
    
    utils.CheckErr(viper.ReadInConfig())
    utils.CheckErr(viper.UnmarshalExact(&configs, func(decoderCfg *mapstructure.DecoderConfig) {
        decoderCfg.WeaklyTypedInput = true
        decoderCfg.ZeroFields = false
    }))
}