package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ul-gaul/stratos_ground-station/backend/acquisition/utils"
	"github.com/ul-gaul/stratos_ground-station/backend/config"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Todo",
		Long:  "TODO",
	}
)

func Execute(fnc func(cmd *cobra.Command, args []string) error) error {
	rootCmd.RunE = fnc
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initialize)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbose output / log level")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("load-csv", "l", "", "Load data from file")
	utils.CheckErr(rootCmd.MarkFlagFilename("load-csv", "csv"))
	utils.CheckErr(rootCmd.MarkPersistentFlagFilename("config", "yml", "yaml"))
}

func initialize() {
	config.InitConfig(cfgFile)
}

func GetCSVFile() string {
	csv, err := rootCmd.Flags().GetString("load-csv")
	if err != nil {
		log.Error(err)
	}
	return csv
}
