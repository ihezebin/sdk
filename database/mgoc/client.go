package mgoc

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/utils/stringer"
	"time"
)

type Client interface {
	Kernel() *mgo.Session
	Close()
	Do(ctx context.Context, model Model, exec func(s *mgo.Collection) error) error
	Config() Config
	NewBaseModel(database string, collection string) *Base
}

// example: mongodb://username:password@localhost:27017,otherhost:27017/db, default auto_time is true
func Dial(url string) (Client, error) {
	info, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}
	config := Config{
		Addrs:          info.Addrs,
		Database:       info.Database,
		Username:       info.Username,
		Password:       info.Password,
		Source:         info.Source,
		ReplicaSetName: info.ReplicaSetName,
		Timeout:        info.Timeout,
		PoolLimit:      info.PoolLimit,
		MaxIdleTime:    time.Duration(info.MaxIdleTimeMS) * time.Millisecond,
		AppName:        info.AppName,
		AutoTime:       true,
	}
	return NewClient(config)
}

func NewClient(config Config) (Client, error) {

	// Determine whether a client with the same alias already exists
	if stringer.NotEmpty(config.Alias) && ClientManager().Has(config.Alias) {
		return nil, errors.Errorf("there is already a client with alias %s, you can choose to use another alias", config.Alias)
	}

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

	// if alias is not empty, add client to the clientMap
	if stringer.NotEmpty(config.Alias) {
		ClientManager().Add(c.config.Alias, c)
	}

	return c, nil
}

type client struct {
	session *mgo.Session
	config  Config
}

func (c *client) NewBaseModel(database string, collection string) *Base {
	return NewBaseModel(c, database, collection)
}

//Deprecated: Use Kernel instead.
func (c *client) Session() *mgo.Session {
	return c.session.Copy()
}

func (c *client) Kernel() *mgo.Session {
	return c.Session()
}

func (c *client) Close() {
	if stringer.NotEmpty(c.Config().Alias) {
		ClientManager().Delete(c.Config().Alias)
	}
	c.session.Close()
}

func (c *client) Config() Config {
	return c.config
}

func (c *client) Do(ctx context.Context, model Model, exec func(c *mgo.Collection) error) error {
	s := c.Kernel()
	defer s.Close()
	return exec(s.DB(model.Database()).C(model.Collection()))
}

func (c *client) DoWithSession(exec func(session *mgo.Session) error) error {
	session := c.Kernel()
	defer session.Close()
	return exec(session)
}

// NewGlobalClient If there is only one Mongo, you can select the global client
func NewGlobalClient(config Config) (Client, error) {
	var err error
	gClient, err = NewClient(config)
	return gClient, err
}

var gClient Client

func GetGlobalClient() Client {
	return gClient
}
