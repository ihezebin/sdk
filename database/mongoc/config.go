package mongoc

import "time"

const (
	defaultPoolLimit   uint64 = 50
	defaultMaxIdleTime        = time.Duration(20) * time.Minute
	defaultSource             = "admin"
)

type Config struct {
	Addrs          []string `json:"addrs"`
	Auth           *Auth    `json:"auth"`
	ReplicaSetName string   `json:"replica_set_name"`
	// Timeout the unit is seconds
	Timeout time.Duration `json:"timeout"`
	// Mode      mgo.Mode      `json:"mode"`
	PoolLimit uint64 `json:"pool_limit"`
	// MaxIdleTime The maximum number of seconds to remain idle in the pool, default 20 minutes
	MaxIdleTime time.Duration `json:"max_idle_time"`
	AppName     string        `json:"app_name"`
	//AutoTime    bool          `json:"auto_time"`
	Alias string `json:"alias"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Source   string `json:"source"`
}
