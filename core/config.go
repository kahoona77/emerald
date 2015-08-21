package core

import (
	"log"
	"os"

	"github.com/kahoona77/emerald/models"

	"code.google.com/p/gcfg"
)

type configFile struct {
	Server models.AppConfig
}

const defaultConfig = `
    [server]
    port = 8080
    logFile = emerald.log
    mongodb = localhost
`

//LoadConfiguration loads the configuration file
func LoadConfiguration(cfgFile string) models.AppConfig {
	var err error
	var cfg configFile

	if cfgFile != "" {
		err = gcfg.ReadFileInto(&cfg, cfgFile)
	} else {
		err = gcfg.ReadStringInto(&cfg, defaultConfig)
	}

	if err != nil {
		panic(err)
	}

	setupLog(cfg.Server.LogFile)

	return cfg.Server
}

func setupLog(logFile string) {
	// setup log
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	// defer f.Close()
	log.SetOutput(f)
}
