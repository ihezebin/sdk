package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk/httpserver/middleware"
)

const ModeDebug = gin.DebugMode
const ModeRelease = gin.ReleaseMode
const ModeTest = gin.TestMode

type Config struct {
	Mode        string `json:"mode"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	middlewares []middleware.Middleware
}

type ConfigFunc func(config *Config)

func newConfig(cfs ...ConfigFunc) Config {
	config := Config{
		Mode:        gin.DebugMode,
		Name:        "",
		Port:        8080,
		middlewares: nil,
	}
	for _, setConfig := range cfs {
		setConfig(&config)
	}
	return config
}

func WithMode(mode string) ConfigFunc {
	return func(config *Config) {
		switch mode {
		case ModeDebug:
			config.Mode = mode
		case ModeRelease:
			config.Mode = mode
		case ModeTest:
			config.Mode = mode
		default:
			panic("available mode: debug release test")
		}
	}
}

func WithName(name string) ConfigFunc {
	return func(config *Config) {
		config.Name = name
	}
}

func WithPort(port int) ConfigFunc {
	return func(config *Config) {
		config.Port = port
	}
}

func WithMiddles(middlewares ...middleware.Middleware) ConfigFunc {
	return func(config *Config) {
		config.middlewares = middlewares
	}
}
