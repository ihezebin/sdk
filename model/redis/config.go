package redis

import (
	"time"
)

const (
	defaultNetwork = "tcp"
)

type Config struct {
	Addr     string
	Username string
	Password string
	// Maximum number of idle connections in the pool.
	MaxIdle int
	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int
	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration
	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool
	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration

	ClientName string
	Database   int
}

//type Model interface {
//	ModelName() string
//}
