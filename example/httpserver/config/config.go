package config

import (
	"github.com/whereabouts/sdk/config"
)

// Config global config
type Config struct {
	Port    int    `json:"port"`
	AppName string `json:"app_name"`
	Mode    string `json:"mode"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}

func Load() error {
	return config.LoadWithDefault(&gConfig)
}
