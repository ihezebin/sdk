package config

// Config global configure
type Config struct {
	Port    int    `json:"port"`
	AppName string `json:"app_name"`
	Mode    string `json:"mode"`
}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}
