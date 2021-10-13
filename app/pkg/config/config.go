package config

import (
	"log"
	"tasks/app/pkg/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	LogFile   string
	Templates string
	SQLDriver string
	DBName    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("./pkg/config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		LogFile:   cfg.Section("web").Key("logFile").String(),
		Templates: cfg.Section("views").Key("templates").String(),
		SQLDriver: cfg.Section("database").Key("driver").String(),
		DBName:    cfg.Section("database").Key("dbname").String(),
	}
}
