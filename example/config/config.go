package main

type Config struct {
	Port    int    `json:"port"`
	Name    string `json:"name"`
	RealAge int    `json:"real_age"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}
