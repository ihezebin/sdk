package mongoc

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (model *Base) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.CountDocuments(ctx, filter, opts...)
	})
	return count.(int64), err
}

func (model *Base) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return model.Count(ctx, filter, opts...)
}

func (model *Base) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.InsertMany(ctx, documents, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.InsertManyResult), err
}

func (model *Base) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.InsertOne(ctx, document, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.InsertOneResult), err
}

func (model *Base) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.DeleteMany(ctx, filter, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.DeleteResult), err
}

func (model *Base) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.DeleteOne(ctx, filter, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.DeleteResult), err
}

func (model *Base) DeleteId(ctx context.Context, id interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return model.DeleteOne(ctx, model.wrapId(id), opts...)
}

func (model *Base) Find(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		cur, err := collection.Find(ctx, filter, opts...)
		if err == nil {
			err = cur.All(ctx, results)
		}
		return nil, err
	})
	return err
}

func (model *Base) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.FindOne(ctx, filter, opts...).Decode(result)
	})
	return err
}

func (model *Base) FindId(ctx context.Context, id interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return model.FindOne(ctx, model.wrapId(id), result, opts...)
}

func (model *Base) FindOneAndDelete(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneAndDeleteOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.FindOneAndDelete(ctx, filter, opts...).Decode(result)
	})
	return err
}

func (model *Base) FindOneAndReplace(ctx context.Context, filter, replacement, result interface{}, opts ...*options.FindOneAndReplaceOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.FindOneAndReplace(ctx, filter, replacement, opts...).Decode(result)
	})
	return err
}

func (model *Base) FindOneAndUpdate(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
	})
	return err
}

func (model *Base) FindOneAndUpsert(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		optUpsert := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
		opts = append(opts, optUpsert)
		return nil, collection.FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
	})
	return err
}

func (model *Base) ReplaceOne(ctx context.Context, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.ReplaceOne(ctx, filter, replacement, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.UpdateResult), err
}

func (model *Base) UpdateMany(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.UpdateMany(ctx, filter, replacement, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.UpdateResult), err
}

func (model *Base) UpdateOne(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.UpdateOne(ctx, filter, replacement, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.UpdateResult), err
}

func (model *Base) UpdateId(ctx context.Context, id, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return model.UpdateOne(ctx, model.wrapId(id), replacement, opts...)
}

func (model *Base) Upsert(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		optUpsert := options.Update().SetUpsert(true)
		opts = append(opts, optUpsert)
		return collection.UpdateOne(ctx, filter, replacement, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.UpdateResult), err
}

func (model *Base) UpsertId(ctx context.Context, id, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return model.Upsert(ctx, model.wrapId(id), replacement, opts...)
}

func (model *Base) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.Distinct(ctx, fieldName, filter, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.([]interface{}), err
}

func (model *Base) Aggregate(ctx context.Context, pipeline, results interface{}, opts ...*options.AggregateOptions) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		cur, err := collection.Aggregate(ctx, pipeline, opts...)
		if err == nil {
			err = cur.All(ctx, results)
		}
		return nil, err
	})
	return err
}

func (model *Base) Indexes() mongo.IndexView {
	indexes, _ := model.Do(context.TODO(), func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.Indexes(), nil
	})
	return indexes.(mongo.IndexView)
}

func (model *Base) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	col, err := model.Do(context.TODO(), func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.Clone(opts...)
	})
	if col == nil {
		return nil, err
	}
	return col.(*mongo.Collection), err
}

func (model *Base) Drop(ctx context.Context) error {
	_, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return nil, collection.Drop(ctx)
	})
	return err
}

func (model *Base) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	result, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.BulkWrite(ctx, models, opts...)
	})
	if result == nil {
		return nil, err
	}
	return result.(*mongo.BulkWriteResult), err
}

func (model *Base) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	count, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.EstimatedDocumentCount(ctx, opts...)
	})
	return count.(int64), err
}

func (model *Base) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	stream, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.Watch(ctx, pipeline, opts...)
	})
	return stream.(*mongo.ChangeStream), err
}

func (model *Base) wrapId(id interface{}) interface{} {
	return bson.M{"_id": id}
}
