package config

import (
	"github.com/ihezebin/sdk/config"
)

// Config global config
type Config struct {
	Port    int    `mapstructure:"port" json:"port"`
	AppName string `mapstructure:"app_name" json:"app_name"`
	Mode    string `mapstructure:"mode" json:"mode"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}

func Load() error {
	return config.NewWithFilePath("./config.toml").Load(&gConfig)
}
