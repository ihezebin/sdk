package mongoc

import (
	"github.com/whereabouts/sdk/utils/stringer"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	defaultPoolLimit   uint64 = 50
	defaultMaxIdleTime        = time.Duration(20) * time.Minute
	defaultSource             = "admin"
)

type Config struct {
	Addrs          []string `mapstructure:"addrs" json:"addrs"`
	Auth           *Auth    `mapstructure:"auth" json:"auth"`
	ReplicaSetName string   `mapstructure:"replica_set_name" json:"replica_set_name"`
	// Timeout the unit is seconds
	Timeout time.Duration `mapstructure:"timeout" json:"timeout"`
	// Mode      mgo.Mode      `mapstructure:"mode"`
	PoolLimit uint64 `mapstructure:"pool_limit" json:"pool_limit"`
	// MaxIdleTime The maximum number of seconds to remain idle in the pool, default 20 minutes
	MaxIdleTime time.Duration `mapstructure:"max_idle_time" json:"max_idle_time"`
	AppName     string        `mapstructure:"app_name" json:"app_name"`
}

type Auth struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Source   string `mapstructure:"source" json:"source"`
}

func convertConfig2Options(config Config) options.ClientOptions {
	opts := options.Client()
	opts.SetHosts(config.Addrs)
	if config.Auth != nil {
		source := defaultSource
		if stringer.NotEmpty(config.Auth.Source) {
			source = config.Auth.Source
		}
		opts.SetAuth(options.Credential{
			Username:   config.Auth.Username,
			Password:   config.Auth.Password,
			AuthSource: source,
		})
	}
	if stringer.NotEmpty(config.ReplicaSetName) {
		opts.SetReplicaSet(config.ReplicaSetName)
	}
	poolLimit := defaultPoolLimit
	if config.PoolLimit != 0 {
		poolLimit = config.PoolLimit
	}
	opts.SetMaxPoolSize(poolLimit)

	maxIdleTime := defaultMaxIdleTime
	if config.MaxIdleTime != time.Duration(0) {
		maxIdleTime = config.MaxIdleTime * time.Second
	}
	opts.SetMaxConnIdleTime(maxIdleTime)
	if config.MaxIdleTime != time.Duration(0) {
		opts.SetSocketTimeout(config.Timeout * time.Second)
	}

	if stringer.NotEmpty(config.AppName) {
		opts.SetAppName(config.AppName)
	}

	return *opts
}
