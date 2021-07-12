package mongoc

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"time"
)

type Client interface {
	GetSession() *mgo.Session
	Close()
	DoWithContext(ctx context.Context, model Model, exec func(s *mgo.Collection) error) error
	GetConfig() Config
}

// example: mongodb://username:password@localhost:27017,otherhost:27017/db
func Dial(url string) (Client, error) {
	info, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}
	if info.PoolLimit == 0 {
		info.PoolLimit = defaultPoolLimit
	}
	if info.MaxIdleTimeMS == 0 {
		info.MaxIdleTimeMS = int(defaultMaxIdleTime / time.Millisecond)
	}
	if info.Source == "" {
		info.Source = "admin"
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect mongodb")
	}
	session.SetMode(mgo.PrimaryPreferred, true)
	config := Config{
		Addrs:          info.Addrs,
		Database:       info.Database,
		Username:       info.Username,
		Password:       info.Password,
		Source:         info.Source,
		ReplicaSetName: info.ReplicaSetName,
		Timeout:        info.Timeout,
		Mode:           session.Mode(),
		PoolLimit:      info.PoolLimit,
		MaxIdleTime:    time.Duration(info.MaxIdleTimeMS) * time.Millisecond,
		AppName:        info.AppName,
		InsertTimeAuto: true,
		UpdateTimeAuto: true,
	}
	c := &client{session: session, config: config}
	gClient = c
	return c, nil
}

func NewClient(config Config) (Client, error) {
	poolLimit := defaultPoolLimit
	if config.PoolLimit != 0 {
		poolLimit = config.PoolLimit
	}

	maxIdleTime := defaultMaxIdleTime
	if config.MaxIdleTime != time.Duration(0) {
		maxIdleTime = config.MaxIdleTime * time.Second
	}

	source := defaultSource
	if config.Source != "" {
		source = config.Source
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:          config.Addrs,
		Database:       config.Database,
		Username:       config.Username,
		Password:       config.Password,
		Source:         source,
		ReplicaSetName: config.ReplicaSetName,
		Timeout:        config.Timeout * time.Second,
		PoolLimit:      poolLimit,
		MaxIdleTimeMS:  int(maxIdleTime / time.Millisecond),
		AppName:        config.AppName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect mongodb")
	}

	if config.Mode < mgo.Primary {
		config.Mode = mgo.PrimaryPreferred
	}
	session.SetMode(config.Mode, true)
	c := &client{session: session, config: config}
	gClient = c
	return c, nil
}

type client struct {
	session *mgo.Session
	config  Config
}

//Deprecated: Use DoWithSession instead.
func (c *client) GetSession() *mgo.Session {
	return c.session.Copy()
}

func (c *client) Close() {
	c.session.Close()
}

func (c *client) DoWithContext(ctx context.Context, model Model, exec func(c *mgo.Collection) error) error {
	s := c.GetSession()
	defer s.Close()
	return exec(s.DB(model.Database()).C(model.Collection()))
}

func (c *client) DoWithSession(do func(session *mgo.Session) error) error {
	session := c.GetSession()
	defer session.Close()
	return do(session)
}

func (c *client) GetConfig() Config {
	return c.config
}

var gClient *client

func getGlobalClient() Client {
	return gClient
}
