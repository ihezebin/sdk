package mongoc

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

const (
	defaultPoolLimit   = 50
	defaultMaxIdleTime = time.Duration(20) * time.Minute
	defaultSource      = "admin"
)

var (
	NullRet = &struct{}{}
	NullDoc = make(bson.M)
	NullMap = make(map[string]interface{})
)

type Config struct {
	Addrs          []string `json:"addrs"`
	Database       string   `json:"database"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Source         string   `json:"source"`
	ReplicaSetName string   `json:"replica_set_name"`
	// Timeout the unit is seconds
	Timeout   time.Duration `json:"timeout"`
	Mode      mgo.Mode      `json:"mode"`
	PoolLimit int           `json:"pool_limit"`
	// MaxIdleTime The maximum number of seconds to remain idle in the pool, default 20 minutes
	MaxIdleTime time.Duration `json:"max_idle_time"`
	AppName     string        `json:"app_name"`
	AutoTime    bool          `json:"auto_time"`
}
