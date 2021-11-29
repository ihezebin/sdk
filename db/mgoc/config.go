package mgoc

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
	Addrs          []string `mapstructure:"addrs" json:"addrs"`
	Database       string   `mapstructure:"db" json:"db"`
	Username       string   `mapstructure:"username" json:"username"`
	Password       string   `mapstructure:"password" json:"password"`
	Source         string   `mapstructure:"source" json:"source"`
	ReplicaSetName string   `mapstructure:"replica_set_name" json:"replica_set_name"`
	// Timeout the unit is seconds
	Timeout   time.Duration `mapstructure:"timeout" json:"timeout"`
	Mode      mgo.Mode      `mapstructure:"mode" json:"mode"`
	PoolLimit int           `mapstructure:"pool_limit" json:"pool_limit"`
	// MaxIdleTime The maximum number of seconds to remain idle in the pool, default 20 minutes
	MaxIdleTime time.Duration `mapstructure:"max_idle_time" json:"max_idle_time"`
	AppName     string        `mapstructure:"app_name" json:"app_name"`
	// Alias If it is empty, it will not be included in the client Map management
	// 如果为空则不纳入clientMap管理
	Alias string `mapstructure:"alias" json:"alias"`
}
