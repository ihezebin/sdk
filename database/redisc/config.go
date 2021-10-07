package redisc

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

	ClientName string `mapstructure:"client_name" json:"client_name"`
	Database   int    `mapstructure:"database" json:"database"`
}

//type Model interface {
//	ModelName() string
//}
