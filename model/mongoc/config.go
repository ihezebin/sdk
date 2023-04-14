package mongoc

import (
	"github.com/ihezebin/sdk/utils/stringer"
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

func (c Config) Convert2Options() *options.ClientOptions {
	opts := options.Client()
	opts.SetHosts(c.Addrs)
	if c.Auth != nil {
		source := defaultSource
		if stringer.NotEmpty(c.Auth.Source) {
			source = c.Auth.Source
		}
		opts.SetAuth(options.Credential{
			Username:   c.Auth.Username,
			Password:   c.Auth.Password,
			AuthSource: source,
		})
	}
	if stringer.NotEmpty(c.ReplicaSetName) {
		opts.SetReplicaSet(c.ReplicaSetName)
	}
	poolLimit := defaultPoolLimit
	if c.PoolLimit != 0 {
		poolLimit = c.PoolLimit
	}
	opts.SetMaxPoolSize(poolLimit)

	maxIdleTime := defaultMaxIdleTime
	if c.MaxIdleTime != time.Duration(0) {
		maxIdleTime = c.MaxIdleTime * time.Second
	}
	opts.SetMaxConnIdleTime(maxIdleTime)
	if c.MaxIdleTime != time.Duration(0) {
		opts.SetSocketTimeout(c.Timeout * time.Second)
	}

	if stringer.NotEmpty(c.AppName) {
		opts.SetAppName(c.AppName)
	}

	return opts
}
