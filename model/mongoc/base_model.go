package mongoc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type baseModel struct {
	*wrapCollection
	database   string
	collection string
	client     *Client
}

func NewBaseModel(client *Client, database string, collection string) *baseModel {
	return &baseModel{
		wrapCollection: client.Kernel().Database(database).Collection(collection),
		database:       database,
		collection:     collection,
		client:         client,
	}
}

func (m *baseModel) Database() string {
	return m.database
}

func (m *baseModel) Collection() string {
	return m.collection
}

func (m *baseModel) Kernel() *mongo.Collection {
	return m.wrapCollection
}

func (m *baseModel) Do(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.Do(ctx, m, exec)
}

func (m *baseModel) DoWithTransaction(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.DoWithTransaction(ctx, m, func(ctx context.Context, model Model) (interface{}, error) {
		return exec(ctx, m)
	})
}

func (m *baseModel) DoWithSession(ctx context.Context, exec func(sessionCtx mongo.SessionContext, model Model) error) error {
	return m.client.DoWithSession(ctx, m, exec)
}

func (m *baseModel) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return m.wrapCollection.FindOne(ctx, filter, opts...).Decode(result)
}

func (m *baseModel) FindWithCursor(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.wrapCollection.Find(ctx, filter, opts...)
}

func (m *baseModel) FindMany(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cursor, err := m.wrapCollection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (m *baseModel) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return m.wrapCollection.CountDocuments(ctx, filter, opts...)
}
