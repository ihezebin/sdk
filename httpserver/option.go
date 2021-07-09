package httpserver

import "github.com/whereabouts/sdk-go/httpserver/middleware"

type Option struct {
	Name        string
	Port        int
	Middlewares []middleware.Middleware
}

type OptionFunc func(option *Option)

func newOption(opts ...OptionFunc) Option {
	option := Option{
		Name:        "",
		Port:        8080,
		Middlewares: nil,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return option
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
