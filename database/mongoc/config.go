package mongoc

import "time"

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
	//AutoTime    bool          `mapstructure:"auto_time"`
	Alias string `mapstructure:"alias" json:"alias"`
}

type Auth struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Source   string `mapstructure:"source" json:"source"`
}
