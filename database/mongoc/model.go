package mongoc

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk/utils/nativer"
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

func (model *Base) Count(ctx context.Context, filter interface{}) (int64, error) {
	count, err := model.Do(ctx, func(ctx context.Context, collection *mongo.Collection) (interface{}, error) {
		return collection.CountDocuments(ctx, filter)
	})
	return count.(int64), err
}

func (model *Base) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := b.c(ctx).CountDocuments(ctx, filter, opts...)
	return count, err
}

func (model *Base) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	insmres, err := b.c(ctx).InsertMany(ctx, documents, opts...)
	return insmres, err
}

func (model *Base) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	insores, err := b.c(ctx).InsertOne(ctx, document, opts...)
	return insores, err
}

func (model *Base) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dmres, err := b.c(ctx).DeleteMany(ctx, filter, opts...)
	return dmres, err
}

func (model *Base) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dor, err := b.c(ctx).DeleteOne(ctx, filter, opts...)
	return dor, err
}

func (model *Base) Find(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cur, err := b.c(ctx).Find(ctx, filter, opts...)
	if err == nil {
		err = cur.All(ctx, results)
	}
	return err
}

func (model *Base) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return b.c(ctx).FindOne(ctx, filter, opts...).Decode(result)
}

func (model *Base) FindOneAndDelete(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneAndDeleteOptions) error {
	return b.c(ctx).FindOneAndDelete(ctx, filter, opts...).Decode(result)
}

func (model *Base) FindOneAndReplace(ctx context.Context, filter, replacement, result interface{}, opts ...*options.FindOneAndReplaceOptions) error {
	return b.c(ctx).FindOneAndReplace(ctx, filter, replacement, opts...).Decode(result)
}

func (model *Base) FindOneAndUpdate(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	return b.c(ctx).FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
}

func (model *Base) FindOneAndUpsert(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	rd := options.After
	optUpsert := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(rd)
	opts = append(opts, optUpsert)
	return b.c(ctx).FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
}

func (model *Base) ReplaceOne(ctx context.Context, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	repres, err := b.c(ctx).ReplaceOne(ctx, filter, replacement, opts...)
	return repres, err
}

func (model *Base) UpdateMany(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	umres, err := b.c(ctx).UpdateMany(ctx, filter, replacement, opts...)
	return umres, err
}

func (model *Base) UpdateOne(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	uores, err := b.c(ctx).UpdateOne(ctx, filter, replacement, opts...)
	return uores, err
}

func (model *Base) Upsert(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	optUpsert := options.Update().SetUpsert(true)
	opts = append(opts, optUpsert)
	uores, err := b.c(ctx).UpdateOne(ctx, filter, replacement, opts...)
	return uores, err
}

func (model *Base) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	distinct, err := b.c(ctx).Distinct(ctx, fieldName, filter, opts...)
	return distinct, err
}

func (model *Base) Aggregate(ctx context.Context, pipeline, results interface{}, opts ...*options.AggregateOptions) error {

	cur, err := b.c(ctx).Aggregate(ctx, pipeline, opts...)
	if err == nil {
		err = cur.All(ctx, results)
	}

	return err
}

func (model *Base) Indexes() mongo.IndexView {
	return b.c(context.Background()).Indexes()
}

func (model *Base) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return b.c(context.Background()).Clone(opts...)
}

func (model *Base) Drop(ctx context.Context) error {
	err := b.c(ctx).Drop(ctx)
	return err
}

func (model *Base) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {

	bwres, err := b.c(ctx).BulkWrite(ctx, models, opts...)
	return bwres, err
}

func (model *Base) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	count, err := b.c(ctx).EstimatedDocumentCount(ctx, opts...)
	return count, err
}

func (model *Base) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	cs, err := b.c(ctx).Watch(ctx, pipeline, opts...)
	return cs, err
}

func (model *Base) wrapId(id interface{}) interface{} {
	return bson.M{"_id": id}
}
