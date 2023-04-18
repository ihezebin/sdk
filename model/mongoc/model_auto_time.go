package mongoc

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

const (
	defaultCreateTimeFieldKey = "create_at"
	defaultUpdateTimeFieldKey = "update_at"
	defaultDeleteTimeFieldKey = "delete_at"
)

type modelAutoTime struct {
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

func NewModelAutoTime(client *Client, database string, collection string) *modelAutoTime {
	return &modelAutoTime{
		wrapCollection:     client.Kernel().Database(database).Collection(collection),
		database:           database,
		collection:         collection,
		client:             client,
		createTimeFieldKey: defaultCreateTimeFieldKey,
		updateTimeFieldKey: defaultUpdateTimeFieldKey,
		deleteTimeFieldKey: defaultDeleteTimeFieldKey,
	}
}

func (m *modelAutoTime) SetSoftDelete(softDelete bool) *modelAutoTime {
	m.softDelete = softDelete
	return m
}

func (m *modelAutoTime) SetCreateTimeFieldKey(createTimeFieldKey string) *modelAutoTime {
	m.createTimeFieldKey = createTimeFieldKey
	return m
}

func (m *modelAutoTime) SetUpdateTimeFieldKey(updateTimeFieldKey string) *modelAutoTime {
	m.updateTimeFieldKey = updateTimeFieldKey
	return m
}

func (m *modelAutoTime) SetDeleteTimeFieldKey(deleteTimeFieldKey string) *modelAutoTime {
	m.deleteTimeFieldKey = deleteTimeFieldKey
	return m
}

func (m *modelAutoTime) Database() string {
	return m.database
}

func (m *modelAutoTime) Collection() string {
	return m.collection
}

func (m *modelAutoTime) Kernel() *mongo.Collection {
	return m.wrapCollection
}

func (m *modelAutoTime) Do(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.Do(ctx, m, exec)
}

func (m *modelAutoTime) DoWithTransaction(ctx context.Context, exec func(ctx context.Context, model Model) (interface{}, error)) (interface{}, error) {
	return m.client.DoWithTransaction(ctx, m, func(ctx context.Context, model Model) (interface{}, error) {
		return exec(ctx, m)
	})
}

func (m *modelAutoTime) DoWithSession(ctx context.Context, exec func(sessionCtx mongo.SessionContext, model Model) error) error {
	return m.client.DoWithSession(ctx, m, exec)
}

func (m *modelAutoTime) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
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

func (m *modelAutoTime) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
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

func (m *modelAutoTime) convert2SoftDeleteFilter(filter interface{}) (interface{}, error) {
	switch v := filter.(type) {
	case bson.M:
		v[m.deleteTimeFieldKey] = bson.M{"$eq": 0}
		return v, nil
	case bson.D:
		v = append(v, bson.E{Key: m.deleteTimeFieldKey, Value: bson.M{"$eq": 0}})
		return v, nil
	case map[string]interface{}:
		v[m.deleteTimeFieldKey] = bson.M{"$eq": 0}
		return v, nil
	default:
		return nil, errors.Errorf("unsupported filter type: %T", filter)
	}
}

func (m *modelAutoTime) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	// 软删除更新删除时间
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
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

func (m *modelAutoTime) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	// 软删除更新删除时间
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
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

func (m *modelAutoTime) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return errors.Wrapf(err, "soft delete filter error")
		}
	}
	return m.wrapCollection.FindOne(ctx, filter, opts...).Decode(result)
}

func (m *modelAutoTime) FindCursor(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
	}
	return m.wrapCollection.Find(ctx, filter, opts...)
}

func (m *modelAutoTime) FindMany(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return errors.Wrapf(err, "soft delete filter error")
		}
	}
	cursor, err := m.wrapCollection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (m *modelAutoTime) convert2Doc(obj interface{}) (bson.M, error) {
	doc := bson.M{}
	objVal := reflect.ValueOf(obj)
	for objVal.Kind() == reflect.Ptr {
		objVal = objVal.Elem()
	}
	switch objVal.Kind() {
	case reflect.Struct:
		for i, objTyp := 0, objVal.Type(); i < objTyp.NumField(); i++ {
			field := objTyp.Field(i)
			if field.Anonymous {
				continue
			}
			tag := field.Tag.Get("bson")
			if tag == "" {
				if tag = field.Tag.Get("json"); tag == "" {
					tag = field.Name
				}
			}
			if tag == "-" {
				continue
			}
			if strings.Contains(tag, ",") {
				tag = strings.Split(tag, ",")[0]
			}
			doc[tag] = objVal.Field(i).Interface()
		}
	case reflect.Map:
		for _, key := range objVal.MapKeys() {
			doc[key.String()] = objVal.MapIndex(key).Interface()
		}
	default:
		return nil, errors.Errorf("unsupported type: %T", obj)
	}

	return doc, nil
}

func (m *modelAutoTime) addUpdateTime2Setter(update interface{}) (interface{}, error) {
	switch v := update.(type) {
	case bson.M: // 通过bson.M来处理setter
		setter, ok := v["$set"]
		if !ok {
			setter = bson.M{}
		}
		setterDoc, err := m.convert2Doc(setter)
		if err != nil {
			return nil, errors.Wrapf(err, "convert update setter to doc error")
		}
		setterDoc[m.updateTimeFieldKey] = time.Now().Unix()
		v["$set"] = setterDoc
		return v, nil
	case bson.D: // bson.D转为bson.M
		updater := bson.M{}
		for _, e := range v {
			updater[e.Key] = e.Value
		}
		return m.addUpdateTime2Setter(updater)
	case map[string]interface{}: // map转为bson.M
		var updater bson.M
		updater = v
		return m.addUpdateTime2Setter(updater)
	default:
		return nil, errors.New("not support update type")
	}
}

func (m *modelAutoTime) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
	}
	var err error
	if update, err = m.addUpdateTime2Setter(update); err != nil {
		return nil, err
	}
	return m.wrapCollection.UpdateOne(ctx, filter, update, opts...)
}

func (m *modelAutoTime) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
	}
	var err error
	if update, err = m.addUpdateTime2Setter(update); err != nil {
		return nil, err
	}
	return m.wrapCollection.UpdateMany(ctx, filter, update, opts...)
}

func (m *modelAutoTime) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return 0, errors.Wrapf(err, "soft delete filter error")
		}
	}
	return m.wrapCollection.CountDocuments(ctx, filter, opts...)
}

func (m *modelAutoTime) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	if m.softDelete {
		var err error
		if filter, err = m.convert2SoftDeleteFilter(filter); err != nil {
			return nil, errors.Wrapf(err, "soft delete filter error")
		}
	}
	return m.wrapCollection.Distinct(ctx, fieldName, filter, opts...)
}
