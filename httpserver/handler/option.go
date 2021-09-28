package handler

type config struct {
	withGinCtx bool
	noResponse bool
}

type Option func(config *config)

func newConfig(options ...Option) config {
	conf := config{}
	for _, option := range options {
		option(&conf)
	}
	return conf
}

func WithGinContext(withGinCtx bool) Option {
	return func(conf *config) {
		conf.withGinCtx = withGinCtx
	}
}

func WithNoResponse(noResp bool) Option {
	return func(conf *config) {
		conf.noResponse = noResp
	}
}
