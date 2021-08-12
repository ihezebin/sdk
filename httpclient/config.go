package httpclient

import "time"

type Config struct {
	Address          string
	Timeout          time.Duration
	RetryCount       int           // 重试次数
	RetryWaitTime    time.Duration // 重试间隔等待时间
	RetryMaxWaitTime time.Duration // 重试间隔最大等待时间
}

type Option func(*Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithAddress(address string) Option {
	return func(config *Config) {
		config.Address = address
	}
}
