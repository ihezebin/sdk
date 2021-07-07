package config

// global configure
type Config struct {
	Port int    `json:"port"`
	Env  string `json:"env"`
}

func (c *Config) Set() {

}

var gConfig Config

func GetConfig() *Config {
	return &gConfig
}
