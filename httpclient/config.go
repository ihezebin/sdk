package httpclient

import "time"

type Config struct {
	Host             string        `mapstructure:"host" json:"host"`
	Timeout          time.Duration `mapstructure:"timeout" json:"timeout"`                         // 单位:s
	RetryCount       int           `mapstructure:"retry_count" json:"retry_count"`                 // 重试次数
	RetryWaitTime    time.Duration `mapstructure:"retry_wait_time" json:"retry_wait_time"`         // 重试间隔等待时间, 单位:s
	RetryMaxWaitTime time.Duration `mapstructure:"retry_max_wait_time" json:"retry_max_wait_time"` // 重试间隔最大等待时间, 单位:s
	Alias            string        `mapstructure:"alias" json:"alias"`
}

type Option func(*Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithHost(host string) Option {
	return func(config *Config) {
		config.Host = host
	}
}

func WithAlias(alias string) Option {
	return func(config *Config) {
		config.Alias = alias
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.Timeout = timeout
	}
}

func WithRetryCount(retryCount int) Option {
	return func(config *Config) {
		config.RetryCount = retryCount
	}
}

func WithRetryWaitTime(retryWaitTime time.Duration) Option {
	return func(config *Config) {
		config.RetryWaitTime = retryWaitTime
	}
}

func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) Option {
	return func(config *Config) {
		config.RetryMaxWaitTime = retryMaxWaitTime
	}
}
