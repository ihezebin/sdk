package v1

import (
	"time"
)

const (
	defaultNetwork = "tcp"
)

type Config struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	// Maximum number of idle connections in the pool.
	MaxIdle int `mapstructure:"max_idle" json:"max_idle"`
	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int `mapstructure:"max_active" json:"max_active"`
	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration `mapstructure:"idle_timeout" json:"idle_timeout"`
	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool `mapstructure:"wait" json:"wait"`
	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime" json:"max_conn_lifetime"`

	// ClientName specifies a client name to be used by the Redis server connection.
	ClientName string `mapstructure:"client_name" json:"client_name"`
	Database   int    `mapstructure:"db" json:"db"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithAddr(addr string) Option {
	return func(config *Config) {
		config.Addr = addr
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

func WithMaxIdle(maxIdle int) Option {
	return func(config *Config) {
		config.MaxIdle = maxIdle
	}
}

func WithMaxActive(maxActive int) Option {
	return func(config *Config) {
		config.MaxActive = maxActive
	}
}

func WithIdleTimeout(idleTimeout time.Duration) Option {
	return func(config *Config) {
		config.IdleTimeout = idleTimeout
	}
}

func WithWait(wait bool) Option {
	return func(config *Config) {
		config.Wait = wait
	}
}

func WithMaxConnLifetime(maxConnLifetime time.Duration) Option {
	return func(config *Config) {
		config.MaxConnLifetime = maxConnLifetime
	}
}

func WithClientName(clientName string) Option {
	return func(config *Config) {
		config.ClientName = clientName
	}
}

func WithDatabase(database int) Option {
	return func(config *Config) {
		config.Database = database
	}
}
