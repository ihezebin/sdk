package main

type Config struct {
	Port int    `json:"port"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}
