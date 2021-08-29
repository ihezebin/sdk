package mongoc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/utils/stringer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Client interface {
	Kernel() *mongo.Client
	Close(ctx context.Context) error
	Config() Config
	Do(ctx context.Context, model Model, exec Exec) (interface{}, error)
	DoWithTransaction(ctx context.Context, model Model, exec Exec) (interface{}, error)
	DoWithSession(ctx context.Context, model Model, exec SessionExec) error
	NewBaseModel(database string, collection string) *Base
}

// NewClient If only one db is used, it is recommended to use: NewGlobalClient
// 若只使用到了一个库，推荐使用: NewGlobalClient
func NewClient(ctx context.Context, config Config) (Client, error) {

	// Determine whether a client with the same alias already exists
	if stringer.NotEmpty(config.Alias) && ClientManager().Has(config.Alias) {
		return nil, errors.Errorf("there is already a client with alias %s, you can choose to use another alias", config.Alias)
	}
	opts := convertConfig2Options(config)

	return NewClientWithOptions(ctx, opts, config.Alias, config.AutoTime)
}

// NewClientWithURI uri format: mongodb://username:password@example1.com,example2.com,example3.com/?replicaSet=test&w=majority&wtimeoutMS=5000
// Alias is used to specify the client alias. When used, the client can be obtained from the clientMap of the manger through the alias, and ClientManager().Get(alias) can be called
// alias用于指定client别名，使用时可以通过该别名从manger的clientMap中获取client, 调用 ClientManager().Get(alias)
// autoTime表示是否开启自动生成插入时间和更新时间
func NewClientWithURI(ctx context.Context, uri string, alias string, autoTime bool) (Client, error) {
	return NewClientWithOptions(ctx, *options.Client().ApplyURI(uri), alias, autoTime)
}

func NewClientWithOptions(ctx context.Context, options options.ClientOptions, alias string, autoTime bool) (Client, error) {
	kernel, err := mongo.Connect(ctx, &options)
	if err != nil {
		return nil, err
	}

	err = kernel.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	config := Config{
		Addrs: options.Hosts,
		Auth: &Auth{
			Username: options.Auth.Username,
			Password: options.Auth.Password,
			Source:   options.Auth.AuthSource,
		},
		AutoTime: autoTime,
		Alias:    alias,
	}

	if options.ReplicaSet != nil {
		config.ReplicaSetName = *options.ReplicaSet
	}
	if options.SocketTimeout != nil {
		config.Timeout = *options.SocketTimeout
	}
	if options.MaxPoolSize != nil {
		config.PoolLimit = *options.MaxPoolSize
	}
	if options.MaxConnIdleTime != nil {
		config.MaxIdleTime = *options.MaxConnIdleTime
	}
	if options.AppName != nil {
		config.AppName = *options.AppName
	}

	c := &client{kernel: kernel, config: config}

	// if alias is not empty, add client to the clientMap
	if stringer.NotEmpty(config.Alias) {
		ClientManager().Add(c.config.Alias, c)
	}

	return c, nil
}

type client struct {
	kernel *mongo.Client
	config Config
}

func (c *client) NewBaseModel(database string, collection string) *Base {
	return NewBaseModel(c, database, collection)
}

func (c *client) Kernel() *mongo.Client {
	return c.kernel
}

func (c *client) Close(ctx context.Context) error {
	if stringer.NotEmpty(c.Config().Alias) {
		ClientManager().Delete(c.Config().Alias)
	}
	return c.Kernel().Disconnect(ctx)
}

func (c *client) Config() Config {
	return c.config
}

type Exec = func(ctx context.Context, collection *mongo.Collection) (interface{}, error)

func (c *client) Do(ctx context.Context, model Model, exec Exec) (interface{}, error) {
	return exec(ctx, c.Kernel().Database(model.Database()).Collection(model.Collection()))
}

type ModelExec = func(ctx context.Context, model *Base) (interface{}, error)

// DoWithTransaction Execute the transaction, if the return err is nil, the transaction is automatically committed, otherwise it is rolled back
// 执行事务, 如果返回err为nil，则事务自动提交, 否则回滚
func (c *client) DoWithTransaction(ctx context.Context, model Model, exec Exec) (interface{}, error) {
	// Specify the DefaultReadConcern option so any transactions started through
	// the session will have read concern majority.
	// The DefaultReadPreference and DefaultWriteConcern options aren't
	// specified so they will be inheritied from client and be set to primary
	// and majority, respectively.
	sessionOpts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	session, err := c.Kernel().StartSession(sessionOpts)
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)
	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	return session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Use sessCtx as the Context parameter for InsertOne and FindOne so
		// both operations are run in a transaction.
		return exec(sessCtx, c.Kernel().Database(model.Database()).Collection(model.Collection()))
	}, txnOpts)
}

type SessionExec = func(sessionCtx mongo.SessionContext, collection *mongo.Collection) error

// DoWithSession To open a transaction in the form of Seesion,
// you need to manually commit the transaction by sessionContext.CommitTransaction(ctx) or rollback the transaction by sessionContext.CommitTransaction(ctx)
// 以Seesion形式开启事务, 需要手动sessionContext.CommitTransaction(ctx)提交事务或sessionContext.CommitTransaction(ctx)回滚事务
func (c *client) DoWithSession(ctx context.Context, model Model, exec SessionExec) error {
	sessionOpts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	session, err := c.Kernel().StartSession(sessionOpts)
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	return mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(); err != nil {
			return err
		}
		return exec(sessionContext, c.Kernel().Database(model.Database()).Collection(model.Collection()))
	})
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
	return *opts
}

// NewGlobalClient If there is only one Mongo, you can select the global client
func NewGlobalClient(ctx context.Context, config Config) (Client, error) {
	var err error
	gClient, err = NewClient(ctx, config)
	return gClient, err
}

var gClient Client

func GetGlobalClient() Client {
	return gClient
}
