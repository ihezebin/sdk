package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/httpserver/middleware"
)

const ModeDebug = gin.DebugMode
const ModeRelease = gin.ReleaseMode
const ModeTest = gin.TestMode

type Option struct {
	Mode        string
	Name        string
	Port        int
	Middlewares []middleware.Middleware
}

type OptionFunc func(option *Option)

func newOption(opts ...OptionFunc) Option {
	option := Option{
		Mode:        gin.DebugMode,
		Name:        "",
		Port:        8080,
		Middlewares: nil,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return option
}

func WithMode(mode string) OptionFunc {
	return func(option *Option) {
		switch mode {
		case ModeDebug:
			option.Mode = mode
		case ModeRelease:
			option.Mode = mode
		case ModeTest:
			option.Mode = mode
		default:
			panic("available mode: debug release test")
		}
	}
}

func WithName(name string) OptionFunc {
	return func(option *Option) {
		option.Name = name
	}
}

func WithPort(port int) OptionFunc {
	return func(option *Option) {
		option.Port = port
	}
}

func WithMiddles(middlewares ...middleware.Middleware) OptionFunc {
	return func(option *Option) {
		option.Middlewares = middlewares
	}
}
