package httpserver

type Option struct {
	Name        string
	Port        string
	Middlewares []Middleware
}

type OptionFunc func(option *Option)

func newOption(opts ...OptionFunc) Option {
	option := Option{
		Name:        "",
		Port:        ":8080",
		Middlewares: nil,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return option
}
