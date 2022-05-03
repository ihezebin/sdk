package mongoc

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	defaultCreateTimeFieldKey = "create_at"
	defaultUpdateTimeFieldKey = "update_at"
	defaultDeleteTimeFieldKey = "delete_at"
)

type autoTimeModel struct {
	*wrapCollection
	database           string
	collection         string
	client             *Client
	softDelete         bool
	createTimeFieldKey string
	updateTimeFieldKey string
	deleteTimeFieldKey string
	millisecond        bool
}

func NewAutoTimeModel(client *Client, database string, collection string) *autoTimeModel {
	return &autoTimeModel{
		wrapCollection:     client.Kernel().Database(database).Collection(collection),
		database:           database,
		collection:         collection,
		client:             client,
		createTimeFieldKey: defaultCreateTimeFieldKey,
		updateTimeFieldKey: defaultUpdateTimeFieldKey,
		deleteTimeFieldKey: defaultDeleteTimeFieldKey,
	}
}

func (m *autoTimeModel) SetSoftDelete(softDelete bool) *autoTimeModel {
	m.softDelete = softDelete
	return m
}

func (m *autoTimeModel) SetCreateTimeFieldKey(createTimeFieldKey string) *autoTimeModel {
	m.createTimeFieldKey = createTimeFieldKey
	return m
}

func (m *autoTimeModel) SetUpdateTimeFieldKey(updateTimeFieldKey string) *autoTimeModel {
	m.updateTimeFieldKey = updateTimeFieldKey
	return m
}

func (m *autoTimeModel) SetDeleteTimeFieldKey(deleteTimeFieldKey string) *autoTimeModel {
	m.deleteTimeFieldKey = deleteTimeFieldKey
	return m
}

func (m *autoTimeModel) Database() string {
	return m.database
}

func (m *autoTimeModel) Collection() string {
	return m.collection
}

func (m *autoTimeModel) Kernel() *mongo.Collection {
	return m.wrapCollection
}

func (m *autoTimeModel) Do(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.Do(ctx, m, exec)
}

func (m *autoTimeModel) DoWithTransaction(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.DoWithTransaction(ctx, m, func(ctx context.Context, model Model) (interface{}, error) {
		return exec(ctx, m)
	})
}

func (m *autoTimeModel) DoWithSession(ctx context.Context, exec func(sessionCtx mongo.SessionContext, model Model) error) error {
	return m.client.DoWithSession(ctx, m, exec)
}

func (m *autoTimeModel) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	doc, err := m.convert2Doc(document)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	doc[m.createTimeFieldKey] = now
	doc[m.updateTimeFieldKey] = now

	if m.softDelete {
		doc[m.deleteTimeFieldKey] = int64(0)
	}

	return m.wrapCollection.InsertOne(ctx, doc, opts...)
}

func (m *autoTimeModel) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	for i := range documents {
		doc, err := m.convert2Doc(documents[i])
		if err != nil {
			return nil, err
		}
		now := time.Now().Unix()
		doc[m.createTimeFieldKey] = now
		doc[m.updateTimeFieldKey] = now

		if m.softDelete {
			doc[m.deleteTimeFieldKey] = int64(0)
		}
		documents[i] = doc
	}

	return m.wrapCollection.InsertMany(ctx, documents, opts...)
}

func (m *autoTimeModel) softDeleteFilter(filter interface{}) interface{} {
	switch v := filter.(type) {
	case bson.M:
		v[m.deleteTimeFieldKey] = bson.M{"$eq": 0}
		filter = v
	case bson.D:
		v = append(v, bson.E{Key: m.deleteTimeFieldKey, Value: bson.M{"$eq": 0}})
	default:
		panic("not support filter type")
	}
	return filter
}

func (m *autoTimeModel) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	// 软删除更新删除时间
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
		now := time.Now().Unix()
		softDelResult, err := m.wrapCollection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				m.deleteTimeFieldKey: now,
				m.updateTimeFieldKey: now,
			},
		})
		if err != nil {
			return nil, err
		}
		return &mongo.DeleteResult{
			DeletedCount: softDelResult.ModifiedCount,
		}, nil
	}
	// 真实删除
	return m.wrapCollection.DeleteOne(ctx, filter, opts...)
}

func (m *autoTimeModel) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	// 软删除更新删除时间
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
		now := time.Now().Unix()
		softDelResult, err := m.wrapCollection.UpdateMany(ctx, filter, bson.M{
			"$set": bson.M{
				m.deleteTimeFieldKey: now,
				m.updateTimeFieldKey: now,
			},
		})
		if err != nil {
			return nil, err
		}
		return &mongo.DeleteResult{
			DeletedCount: softDelResult.ModifiedCount,
		}, nil
	}
	// 真实删除
	return m.wrapCollection.DeleteMany(ctx, filter, opts...)
}

func (m *autoTimeModel) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	return m.wrapCollection.FindOne(ctx, filter, opts...).Decode(result)
}

func (m *autoTimeModel) FindWithCursor(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	return m.wrapCollection.Find(ctx, filter, opts...)
}

func (m *autoTimeModel) FindMany(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	cursor, err := m.wrapCollection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (m *autoTimeModel) convert2Doc(obj interface{}) (bson.M, error) {
	doc := bson.M{}
	data, err := bson.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err = bson.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (m *autoTimeModel) addTime2UpdateSet(update interface{}) (interface{}, error) {
	switch t := update.(type) {
	case bson.M:
		if set, ok := t["$set"]; ok {
			setter, err := m.convert2Doc(set)
			if err != nil {
				return nil, err
			}
			setter[m.updateTimeFieldKey] = time.Now().Unix()
			t["$set"] = setter
			update = t
		}
	case bson.D:
		updater := bson.D{}
		for _, v := range t {
			if v.Key == "$set" {
				setter, err := m.convert2Doc(v.Value)
				if err != nil {
					return nil, err
				}
				setter[m.updateTimeFieldKey] = time.Now().Unix()
				v.Value = setter
			}
			updater = append(updater, v)
			update = updater
		}
	default:
		return nil, errors.New("not support update type")
	}

	return update, nil
}

func (m *autoTimeModel) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	var err error
	if update, err = m.addTime2UpdateSet(update); err != nil {
		return nil, err
	}
	return m.wrapCollection.UpdateOne(ctx, filter, update, opts...)
}

func (m *autoTimeModel) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	var err error
	if update, err = m.addTime2UpdateSet(update); err != nil {
		return nil, err
	}
	return m.wrapCollection.UpdateMany(ctx, filter, update, opts...)
}

func (m *autoTimeModel) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	return m.wrapCollection.CountDocuments(ctx, filter, opts...)
}

func (m *autoTimeModel) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	if m.softDelete {
		filter = m.softDeleteFilter(filter)
	}
	return m.wrapCollection.Distinct(ctx, fieldName, filter, opts...)
}
