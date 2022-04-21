package mongoc

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	Database() string
	Collection() string
}

type wrapCollection = mongo.Collection

type model struct {
	*wrapCollection
	database   string
	collection string
	client     *Client
}

var FilterAll = bson.M{}

func NewModel(client *Client, database string, collection string) *model {
	c := client.Kernel().Database(database).Collection(collection)
	return &model{wrapCollection: c, client: client, database: database, collection: collection}
}

func (m *model) Database() string {
	return m.database
}

func (m *model) Collection() string {
	return m.collection
}

func (m *model) Kernel() *mongo.Collection {
	return m.wrapCollection
}

func (m *model) Do(ctx context.Context, exec Exec) (interface{}, error) {
	return m.client.Do(ctx, m, exec)
}

func (m *model) DoWithTransaction(ctx context.Context, exec Exec) (interface{}, error) {
	return m.client.DoWithTransaction(ctx, m, func(ctx context.Context, _ *mongo.Collection) (interface{}, error) {
		return exec(ctx, m.wrapCollection)
	})
}

func (m *model) DoWithSession(ctx context.Context, exec SessionExec) error {
	return m.client.DoWithSession(ctx, m, exec)
}

func (m *model) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := m.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.CountDocuments(ctx, filter, opts...)
	})
	return count.(int64), err
}

func (m *model) FindWithResults(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	_, err := m.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		cur, err := collection.Find(ctx, filter, opts...)
		if err == nil {
			err = cur.All(ctx, results)
		}
		return nil, err
	})
	return err
}

func (m *model) FindOneWithResult(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	_, err := m.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.FindOne(ctx, filter, opts...).Decode(result)
	})
	return err
}

func (m *model) AggregateWithResults(ctx context.Context, pipeline, results interface{}, opts ...*options.AggregateOptions) error {
	_, err := m.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		cur, err := collection.Aggregate(ctx, pipeline, opts...)
		if err == nil {
			err = cur.All(ctx, results)
		}
		return nil, err
	})
	return err
}
