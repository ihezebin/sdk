package redisc

import (
	"github.com/go-redis/redis/v8"
	"github.com/whereabouts/sdk/utils/stringer"
	"time"
)

type Config struct {
	Addrs    []string `mapstructure:"addrs" json:"addrs"`
	Username string   `mapstructure:"username" json:"username"`
	Password string   `mapstructure:"password" json:"password"`
	Database int      `mapstructure:"db" json:"db"`

	// 连接超时秒数
	// Default is 5 seconds.
	DialTimeout int `mapstructure:"dial_timeout" json:"dial_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout int `mapstructure:"read_timeout" json:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout int `mapstructure:"write_timeout" json:"write_timeout"`

	// 连接池容最大量
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize int `mapstructure:"pool_size" json:"pool_size"`
	// 连接池保持连接状态的最小可用连接数
	// because new connection is slow.
	MinIdleConns int `mapstructure:"min_idle_conns" json:"min_idle_conns"`
	// 连接池全部连接都被占用时的超时时间（秒）
	// Default is ReadTimeout + 1 second.
	PoolTimeout int `mapstructure:"pool_timeout" json:"pool_timeout"`
	// 关闭空闲连接的时间（秒）
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout int `mapstructure:"idle_timeout" json:"idle_timeout"`
}

func (c *Config) Convert2Otions() *redis.UniversalOptions {
	options := &redis.UniversalOptions{}
	if c == nil {
		return options
	}

	if len(c.Addrs) > 0 {
		options.Addrs = c.Addrs
	}
	if stringer.NotEmpty(c.Username) {
		options.Username = c.Username
	}
	if stringer.NotEmpty(c.Password) {
		options.Password = c.Password
	}
	if c.Database != 0 {
		options.DB = c.Database
	}
	if c.DialTimeout != 0 {
		options.DialTimeout = time.Duration(c.DialTimeout) * time.Second
	}
	if c.ReadTimeout != 0 {
		options.ReadTimeout = time.Duration(c.ReadTimeout) * time.Second
	}
	if c.WriteTimeout != 0 {
		options.WriteTimeout = time.Duration(c.WriteTimeout) * time.Second
	}
	if c.PoolSize != 0 {
		options.PoolSize = c.PoolSize
	}
	if c.MinIdleConns != 0 {
		options.MinIdleConns = c.MinIdleConns
	}
	if c.PoolTimeout != 0 {
		options.PoolTimeout = time.Duration(c.PoolTimeout) * time.Second
	}
	if c.IdleTimeout != 0 {
		options.IdleTimeout = time.Duration(c.IdleTimeout) * time.Second
	}

	return options
}

type Option func(config *Config)

func newConfig(options ...Option) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithAddrs(addrs ...string) Option {
	return func(config *Config) {
		config.Addrs = addrs
	}
}

func WithUsername(username string) Option {
	return func(config *Config) {
		config.Username = username
	}
}

func WithPassword(password string) Option {
	return func(config *Config) {
		config.Password = password
	}
}

func WithDatabase(database int) Option {
	return func(config *Config) {
		config.Database = database
	}
}

func WithDialTimeout(dialTimeout int) Option {
	return func(config *Config) {
		config.DialTimeout = dialTimeout
	}
}

func WithReadTimeout(readTimeout int) Option {
	return func(config *Config) {
		config.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout int) Option {
	return func(config *Config) {
		config.WriteTimeout = writeTimeout
	}
}

func WithPoolSize(poolSize int) Option {
	return func(config *Config) {
		config.PoolSize = poolSize
	}
}

func WithMinIdleConns(minIdleConns int) Option {
	return func(config *Config) {
		config.MinIdleConns = minIdleConns
	}
}

func WithPoolTimeout(poolTimeout int) Option {
	return func(config *Config) {
		config.PoolTimeout = poolTimeout
	}
}

func WithIdleTimeout(idleTimeout int) Option {
	return func(config *Config) {
		config.IdleTimeout = idleTimeout
	}
}
