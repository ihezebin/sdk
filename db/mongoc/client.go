package mongoc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type wrapClient = mongo.Client

type Client struct {
	*wrapClient
}

//type Client interface {
//	Kernel() *mongo.Client
//	Close(ctx context.Context) error
//	Config() Config
//	Do(ctx context.Context, model Model, exec Exec) (interface{}, error)
//	DoWithTransaction(ctx context.Context, model Model, exec Exec) (interface{}, error)
//	DoWithSession(ctx context.Context, model Model, exec SessionExec) error
//	NewModel(database string, collection string) Model
//}

// NewClient If only one db is used, it is recommended to use: NewGlobalClient
// 若只使用到了一个库，推荐使用: NewGlobalClient
func NewClient(ctx context.Context, config *Config) (*Client, error) {
	opts := config.Convert2Options()
	return NewClientWithOptions(ctx, opts)
}

// NewClientWithURI uri format: mongodb://username:password@example1.com,example2.com,example3.com/?replicaSet=test&w=majority&wtimeoutMS=5000
func NewClientWithURI(ctx context.Context, uri string) (*Client, error) {
	return NewClientWithOptions(ctx, options.Client().ApplyURI(uri))
}

func NewClientWithOptions(ctx context.Context, opts ...*options.ClientOptions) (*Client, error) {
	client, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Client{wrapClient: client}, nil
}

var gClient *Client

// NewGlobalClient If there is only one Mongo, you can select the global client
func NewGlobalClient(ctx context.Context, config *Config) (*Client, error) {
	var err error
	if gClient, err = NewClient(ctx, config); err != nil {
		return nil, err
	}
	return gClient, nil
}

func GetGlobalClient() *Client {
	return gClient
}

func (c *Client) Kernel() *mongo.Client {
	return c.wrapClient
}

func (c *Client) Do(ctx context.Context, model Model, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return exec(ctx, model)
}

// DoWithTransaction Execute the transaction, if the return err is nil, the transaction is automatically committed, otherwise it is rolled back
// 执行事务, 如果返回err为nil，则事务自动提交, 否则回滚
func (c *Client) DoWithTransaction(ctx context.Context, model Model, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
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
		return exec(sessCtx, model)
	}, txnOpts)
}

// DoWithSession To open a transaction in the form of Seesion,
// you need to manually commit the transaction by sessionContext.CommitTransaction(ctx) or rollback the transaction by sessionContext.CommitTransaction(ctx)
// 以Seesion形式开启事务, 需要手动sessionContext.CommitTransaction(ctx)提交事务或sessionContext.CommitTransaction(ctx)回滚事务
func (c *Client) DoWithSession(ctx context.Context, model Model, exec func(sessionCtx mongo.SessionContext, model Model) error) error {
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
		return exec(sessionContext, model)
	})
}
