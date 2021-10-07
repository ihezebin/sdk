package main

type Config struct {
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
	Age  int    `mapstructure:"age" json:"age"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}
