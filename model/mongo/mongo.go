package mongo

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk-go/utils/mapper"
	"github.com/whereabouts/sdk-go/utils/timer"
	"reflect"
	"time"
)

type MongoDB struct {
	database   string
	collection string
	value      interface{}
}

func New(database string, collection string, value interface{}) *MongoDB {
	return &MongoDB{database: database, collection: collection, value: value}
}

func (db *MongoDB) Database() string {
	return db.database
}

func (db *MongoDB) Collection() string {
	return db.collection
}

func (db *MongoDB) Client() Client {
	return getGlobalClient()
}

// Do it is used for you to use the native mgo interface according to your own needs,
// Use when you can't find the method you want in this package
func (db *MongoDB) Do(f func(c *mgo.Collection) error) error {
	return db.Client().Do(db, f)
}

func (db *MongoDB) Remove(selector interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		return c.Remove(selector)
	})
}

func (db *MongoDB) RemoveID(id interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		return c.RemoveId(id)
	})
}

func (db *MongoDB) RemoveAll(selector interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(func(c *mgo.Collection) error {
		changeInfo, err = c.RemoveAll(selector)
		return err
	})
	return changeInfo, err
}

func (db *MongoDB) handleTimeAuto(doc interface{}, isInsert bool) (map[string]interface{}, error) {
	if doc == nil {
		return nil, nil
	}
	v := reflect.ValueOf(doc)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	now := timer.Format(time.Now())
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
		return nil, nil
	}
	if db.Client().GetConfig().UpdateTimeAuto && !v.MapIndex(reflect.ValueOf("update_time")).IsValid() {
		v.SetMapIndex(reflect.ValueOf("update_time"), reflect.ValueOf(now))
	}
	if db.Client().GetConfig().InsertTimeAuto && !v.MapIndex(reflect.ValueOf("create_time")).IsValid() && isInsert {
		v.SetMapIndex(reflect.ValueOf("create_time"), reflect.ValueOf(now))
	}
	ret := make(map[string]interface{})
	for _, key := range v.MapKeys() {
		ret[key.String()] = v.MapIndex(key).Interface()
	}
	return ret, nil
}

func (db *MongoDB) Insert(doc ...interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		out := make([]interface{}, 0, len(doc))
		for _, in := range doc {
			v, err := db.handleTimeAuto(in, true)
			if err != nil {
				return err
			}
			delete(v, "_id")
			out = append(out, v)
		}
		return c.Insert(out...)
	})
}

// Replace replace the original document as a whole,
// but the value of create_time is the value of the old document
func (db *MongoDB) Replace(selector, update interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		newDoc, err := db.handleTimeAuto(update, true)
		if err != nil {
			return err
		}
		delete(newDoc, "_id")
		oldDoc := make(map[string]interface{})
		err = db.FindOne(selector, nil, &oldDoc)
		if err != nil {
			return errors.New(fmt.Sprintf("do not find the old doc by this selector %+v", err))
		}
		if createTime, ok := oldDoc["create_time"]; ok {
			newDoc["create_time"] = createTime
		}
		err = c.Update(selector, newDoc)
		return err
	})
}

func (db *MongoDB) ReplaceId(id, update interface{}) error {
	return db.Replace(bson.M{"_id": id}, update)
}

func (db *MongoDB) ReplaceAll(selector, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(func(c *mgo.Collection) error {
		var newDoc map[string]interface{}
		newDoc, err = db.handleTimeAuto(update, true)
		if err != nil {
			return err
		}
		oldDoc := make(map[string]interface{})
		err = db.FindOne(selector, nil, &oldDoc)
		if err != nil {
			return errors.New(fmt.Sprintf("do not find old doc by this selector %+v", err))
		}
		if createTime, ok := oldDoc["create_time"]; ok {
			newDoc["create_time"] = createTime
		}
		changeInfo, err = c.UpdateAll(selector, newDoc)
		return err
	})
	return changeInfo, err
}

func (db *MongoDB) Modify(selector, doc bson.M, ret interface{}, deletion ...bool) error {
	if len(doc) == 0 {
		return errors.New(fmt.Sprint("doc does not allow nil!"))
	}
	if ret == nil {
		ret = NullRet
	}
	err := db.Do(func(c *mgo.Collection) error {
		setting := "$unset"
		if len(deletion) == 0 || !deletion[0] {
			setting = "$set"
		}
		v, err := db.handleTimeAuto(mapper.BsonM2Map(doc), false)
		if err != nil {
			return err
		}
		return c.Update(selector, bson.M{setting: v})
	})
	if err != nil {
		return err
	}
	return db.FindOne(selector, nil, ret)
}

func (db *MongoDB) ModifyId(id interface{}, doc bson.M, ret interface{}, deletion ...bool) error {
	return db.Modify(bson.M{"_id": id}, doc, ret, deletion...)
}

func (db *MongoDB) ModifyAll(selector, doc bson.M, deletion ...bool) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(func(c *mgo.Collection) error {
		setting := "$unset"
		if len(deletion) == 0 || !deletion[0] {
			setting = "$set"
		}
		changeInfo, err = c.UpdateAll(selector, bson.M{setting: doc})
		return err
	})
	return changeInfo, err
}

//func (db *MongoDB) Upsert(selector, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
//	err = db.Do(func(c *mgo.Collection) error {
//		changeInfo, err = c.Upsert(selector, update)
//		return err
//	})
//	return changeInfo, err
//}

//func (db *MongoDB) UpsertId(id, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
//	err = db.Do(func(c *mgo.Collection) error {
//		changeInfo, err = c.UpsertId(id, update)
//		return err
//	})
//	return changeInfo, err
//}

func (db *MongoDB) handlePicker(picker []string) interface{} {
	ret := make(bson.M)
	for _, field := range picker {
		ret[field] = 1
	}
	return ret
}

// FindOne the param picker([]string) represents the field to return
func (db *MongoDB) FindOne(selector interface{}, picker []string, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		query := c.Find(selector)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindId(id interface{}, picker []string, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		query := c.FindId(id)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindObjectId(id string, picker []string, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		_id := bson.ObjectIdHex(id)
		query := c.FindId(_id)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(ret)
		return err
	})
}

func (db *MongoDB) FindOneWithSort(selector interface{}, sort []string, picker []string, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
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

func (db *MongoDB) Count(selector interface{}) (count int, err error) {
	err = db.Do(func(c *mgo.Collection) error {
		count, err = c.Find(selector).Count()
		return err
	})
	return count, err
}

func (db *MongoDB) FindAll(selector interface{}, sort []string, picker []string, skip int, limit int, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
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

func (db *MongoDB) PipeAll(selector interface{}, ret interface{}) error {
	return db.Do(func(c *mgo.Collection) error {
		return c.Pipe(selector).All(ret)
	})
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
func (db *MongoDB) Distinct(selector interface{}, key string) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	err := db.Do(func(c *mgo.Collection) error {
		return c.Find(selector).Distinct(key, &ret)
	})

	return ret, err
}
