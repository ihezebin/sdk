package mongoc

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	Kernel() *mongo.Collection
	Database() string
	Collection() string
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error
	FindMany(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error)
	Do(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error)
	DoWithTransaction(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error)
	DoWithSession(ctx context.Context, exec func(sessionCtx mongo.SessionContext, model Model) error) error
	FindWithCursor(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

//type Exec = func(ctx context.Context, model Model) (interface{}, error)

//type SessionExec = func(sessionCtx mongo.SessionContext, model Model) error

type wrapCollection = mongo.Collection

var FilterAll = bson.M{}
