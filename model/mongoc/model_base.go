package mongoc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type modelBase struct {
	*wrapCollection
	database   string
	collection string
	client     *Client
}

func NewModelBase(client *Client, database string, collection string) *modelBase {
	return &modelBase{
		wrapCollection: client.Kernel().Database(database).Collection(collection),
		database:       database,
		collection:     collection,
		client:         client,
	}
}

func (m *modelBase) Database() string {
	return m.database
}

func (m *modelBase) Collection() string {
	return m.collection
}

func (m *modelBase) Kernel() *mongo.Collection {
	return m.wrapCollection
}

func (m *modelBase) Do(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.Do(ctx, m, exec)
}

func (m *modelBase) DoWithTransaction(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.DoWithTransaction(ctx, m, func(ctx context.Context, model Model) (interface{}, error) {
		return exec(ctx, m)
	})
}

func (m *modelBase) DoWithSession(ctx context.Context, exec func(sessionCtx mongo.SessionContext, model Model) error) error {
	return m.client.DoWithSession(ctx, m, exec)
}

func (m *modelBase) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return m.wrapCollection.FindOne(ctx, filter, opts...).Decode(result)
}

func (m *modelBase) FindCursor(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.wrapCollection.Find(ctx, filter, opts...)
}

func (m *modelBase) FindMany(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cursor, err := m.wrapCollection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (m *modelBase) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return m.wrapCollection.CountDocuments(ctx, filter, opts...)
}
