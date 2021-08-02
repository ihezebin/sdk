package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Database() string
	Collection() string
}

type Base struct {
	database   string
	collection string
	ctx        context.Context
	client     Client
}

func NewBaseModel(client Client, database string, collection string) *Base {
	return &Base{client: client, database: database, collection: collection}
}

func (model *Base) Database() string {
	return model.database
}

func (model *Base) Collection() string {
	return model.collection
}

func (model *Base) Client() Client {
	return model.client
}

func (model *Base) Do(ctx context.Context, exec Exec) (interface{}, error) {
	return model.Client().Do(ctx, model, exec)
}

func (model *Base) DoWithTransaction(ctx context.Context, exec Exec) (interface{}, error) {
	return model.Client().DoWithTransaction(ctx, model, exec)
}

func (model *Base) Insert(ctx context.Context, documents ...interface{}) (interface{}, error) {
	return model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		result, err := collection.InsertMany(ctx, documents)
		if result != nil {
			if len(result.InsertedIDs) == 1 {
				return result.InsertedIDs[0], err
			}
			return result.InsertedIDs, err
		}
		return nil, err
	})
}

func (model *Base) DeleteOne(ctx context.Context, documents ...interface{}) (interface{}, error) {
	return model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		result, err := collection.InsertMany(ctx, documents)
		if result != nil {
			return result.InsertedIDs, err
		}
		return nil, err
	})
}

//
//func (model *Base) Find(ctx context.Context, selector interface{}, results interface{}) (interface{}, error) {
//	return model.Client().DoWithTransaction(ctx, model, exec)
//}
