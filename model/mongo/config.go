package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

const (
	defaultPoolLimit   = 50
	defaultMaxIdleTime = time.Duration(20) * time.Minute
)

var (
	NullRet = &struct{}{}
	NullDoc = make(bson.M)
	NullMap = make(map[string]interface{})
)

type Config struct {
	Addrs          []string      `json:"addrs"`
	Database       string        `json:"database"`
	Username       string        `json:"username"`
	Password       string        `json:"password"`
	Source         string        `json:"source"`
	ReplicaSetName string        `json:"replica_set_name"`
	Timeout        time.Duration `json:"timeout"`
	Mode           mgo.Mode      `json:"mode"`
	PoolLimit      int           `json:"pool_limit"`
	MaxIdleTime    time.Duration `json:"max_idle_time"`
	AppName        string        `json:"app_name"`
	InsertTimeAuto bool          `json:"insert_time_auto"`
	UpdateTimeAuto bool          `json:"update_time_auto"`
}

type Model interface {
	Database() string
	Collection() string
}
