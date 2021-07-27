package mongoc

import (
	"context"
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk/utils/mapper"
	"github.com/whereabouts/sdk/utils/timer"
	"reflect"
)

type MongoDB struct {
	database   string
	collection string
	ctx        context.Context
	client     Client
}

func NewMongoDB(client Client, database string, collection string) *MongoDB {
	return &MongoDB{client: client, database: database, collection: collection}
}

func (db *MongoDB) Database() string {
	return db.database
}

func (db *MongoDB) Indexes() ([]mgo.Index, error) {
	return db.Client().GetSession().DB(db.Database()).C(db.Collection()).Indexes()
}

func (db *MongoDB) Collection() string {
	return db.collection
}

func (db *MongoDB) Client() Client {
	return db.client
}

// DoWithContext it is used for you to use the native mgo interface according to your own needs,
// Use when you can't find the method you want in this package
func (db *MongoDB) DoWithContext(ctx context.Context, f func(c *mgo.Collection) error) error {
	return db.Client().DoWithContext(ctx, db, f)
}

func (db *MongoDB) RemoveOne(ctx context.Context, selector interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.Remove(selector)
	})
}
func (db *MongoDB) RemoveId(ctx context.Context, id interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.RemoveId(id)
	})
}
func (db *MongoDB) RemoveAll(ctx context.Context, selector interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.RemoveAll(selector)
		return err
	})
	return changeInfo, err
}

func (db *MongoDB) DeleteOne(ctx context.Context, selector interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.Remove(selector)
	})
}

func (db *MongoDB) DeleteId(ctx context.Context, id interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.RemoveId(id)
	})
}

func (db *MongoDB) DeleteAll(ctx context.Context, selector interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.RemoveAll(selector)
		return err
	})
	return changeInfo, err
}

func (db *MongoDB) Insert(ctx context.Context, doc ...interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		out := make([]interface{}, 0, len(doc))
		for _, in := range doc {
			v, err := db.handleTimeAuto(in, true)
			if err != nil {
				return err
			}
			out = append(out, v)
		}
		return c.Insert(out...)
	})
}

// ReplaceOne replace the original document as a whole,
// but the value of create_time is the value of the old document
func (db *MongoDB) ReplaceOne(ctx context.Context, selector, doc interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		newDoc, err := db.handleTimeAuto(doc, false)
		if err != nil {
			return err
		}
		oldDoc := make(map[string]interface{})
		err = db.FindOne(ctx, selector, nil, &oldDoc)
		if err != nil {
			return errors.New(fmt.Sprintf("do not find the old doc by this selector %+v", err))
		}
		// keep the create time
		if createTime, ok := oldDoc[keyCreateTime]; ok {
			newDoc[keyCreateTime] = createTime
		}
		err = c.Update(selector, newDoc)
		return err
	})
}

func (db *MongoDB) ReplaceId(ctx context.Context, id, doc interface{}) error {
	return db.ReplaceOne(ctx, bson.M{"_id": id}, doc)
}

func (db *MongoDB) ModifyOne(ctx context.Context, selector, update interface{}, deletion ...bool) error {
	//if doc == nil {
	//	return errors.New(fmt.Sprint("doc does not allow nil!"))
	//}
	err := db.DoWithContext(ctx, func(c *mgo.Collection) error {
		setType := "$unset"
		if len(deletion) == 0 || !deletion[0] {
			setType = "$set"
		}
		v, err := db.handleTimeAuto(update, false)
		if err != nil {
			return err
		}
		return c.Update(selector, bson.M{setType: v})
	})
	return err
}

func (db *MongoDB) ModifyId(ctx context.Context, id, update interface{}, deletion ...bool) error {
	return db.ModifyOne(ctx, bson.M{"_id": id}, update, deletion...)
}

func (db *MongoDB) ModifyAll(ctx context.Context, selector, update interface{}, deletion ...bool) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		setType := "$set"
		if len(deletion) > 0 && deletion[0] {
			setType = "$unset"
		}
		v, errt := db.handleTimeAuto(update, false)
		if errt != nil {
			return err
		}
		changeInfo, err = c.UpdateAll(selector, bson.M{setType: v})
		return err
	})
	return
}

func (db *MongoDB) Upsert(ctx context.Context, selector, doc interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		out, errT := db.handleTimeAuto(doc, true)
		if errT != nil {
			return errT
		}
		changeInfo, err = c.Upsert(selector, out)
		return err
	})
	return
}

func (db *MongoDB) UpsertId(ctx context.Context, id, doc interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		out, errT := db.handleTimeAuto(doc, true)
		if errT != nil {
			return errT
		}
		changeInfo, err = c.UpsertId(id, out)
		return err
	})
	return
}

// FindOne the param picker([]string) represents the field to return
func (db *MongoDB) FindOne(ctx context.Context, selector interface{}, picker []string, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		query := c.Find(selector)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindId(ctx context.Context, id interface{}, picker []string, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		query := c.FindId(id)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindAll(ctx context.Context, selector interface{}, sort []string, picker []string, skip int, limit int, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		query := c.Find(selector)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		if sort != nil {
			query = query.Sort(sort...)
		}
		if skip > 0 {
			query.Skip(skip)
		}
		if limit > 0 {
			query.Limit(limit)
		}
		return query.All(ret)
	})
}

func (db *MongoDB) FindObjectIdHex(ctx context.Context, id string, picker []string, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		_id := bson.ObjectIdHex(id)
		query := c.FindId(_id)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindOneWithSort(ctx context.Context, selector interface{}, sort []string, picker []string, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		query := c.Find(selector)
		if sort != nil {
			query = query.Sort(sort...)
		}
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) Count(ctx context.Context, selector interface{}) (count int, err error) {
	err = db.DoWithContext(ctx, func(c *mgo.Collection) error {
		count, err = c.Find(selector).Count()
		return err
	})
	return count, err
}

// Distinct unmarshals into result the list of distinct values for the given key.
//
// For example:
// 		ret, err = db.Distinct(bson.M{"gender": 1}, "age")
// 		fmt.Println(ret)
//
// DB:
// 		{ ObjectId("603a081694ea2e906792a8f1"), name:"a", gender:"1", age:12 }
// 		{ ObjectId("603a081694ea2e906792a8f2"), name:"b", gender:"1", age:13 }
// 		{ ObjectId("603a081694ea2e906792a8f3"), name:"c", gender:"1", age:14 }
// 		{ ObjectId("603a081694ea2e906792a8f4"), name:"d", gender:"1", age:15 }
// 		{ ObjectId("603a081694ea2e906792a8f5"), name:"e", gender:"1", age:14 }
//  	{ ObjectId("603a081694ea2e906792a8f6"), name:"f", gender:"1", age:13 }
//
// Console:
// 		[12, 13, 14 ,15]
//
func (db *MongoDB) Distinct(ctx context.Context, selector interface{}, key string) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	err := db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.Find(selector).Distinct(key, &ret)
	})
	return ret, err
}

func (db *MongoDB) PipeAll(ctx context.Context, pipeline interface{}, ret interface{}) error {
	return db.DoWithContext(ctx, func(c *mgo.Collection) error {
		return c.Pipe(pipeline).All(ret)
	})
}

func (db *MongoDB) handlePicker(picker []string) interface{} {
	ret := make(bson.M)
	for _, field := range picker {
		ret[field] = 1
	}
	return ret
}

func (db *MongoDB) handleTimeAuto(doc interface{}, isInsert bool) (map[string]interface{}, error) {
	if doc == nil {
		return nil, errors.New("the doc can not be nil")
	}
	v := reflect.ValueOf(doc)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		m, err := mapper.Struct2Map(v.Interface())
		if err != nil {
			return nil, err
		}
		v = reflect.ValueOf(m)
	}
	if v.Kind() != reflect.Map {
		return nil, errors.New(fmt.Sprintf("the doc %+v is not a map or struct", v.Interface()))
	}
	if len(v.MapKeys()) == 0 {
		return nil, errors.New("the doc can not be no field")
	}

	if db.Client().GetConfig().AutoTime {
		if v.MapIndex(reflect.ValueOf(keyCreateTime)).IsValid() {
			if _, ok := v.MapIndex(reflect.ValueOf(keyCreateTime)).Interface().(string); !ok {
				return nil, errors.New("if the auto_time is to true, that create time must be of type string")
			}
		}
		if v.MapIndex(reflect.ValueOf(keyUpdateTime)).IsValid() {
			if _, ok := v.MapIndex(reflect.ValueOf(keyUpdateTime)).Interface().(string); !ok {
				return nil, errors.New("if the auto_time is to true, that update time must be of type string")
			}
		}
		now := timer.Now()
		if isInsert {
			v.SetMapIndex(reflect.ValueOf(keyCreateTime), reflect.ValueOf(now))
		}
		v.SetMapIndex(reflect.ValueOf(keyUpdateTime), reflect.ValueOf(now))
	}
	ret := make(map[string]interface{})
	for _, key := range v.MapKeys() {
		ret[key.String()] = v.MapIndex(key).Interface()
	}
	delete(ret, "_id")
	return ret, nil
}

const (
	DefaultKeyCreateTime = "create_time"
	DefaultKeyUpdateTime = "update_time"
)

var (
	keyCreateTime = DefaultKeyCreateTime
	keyUpdateTime = DefaultKeyUpdateTime
)

func SetKeyCreateTime(key string) {
	keyCreateTime = key
}

func SetKeyUpdateTime(key string) {
	keyUpdateTime = key
}
