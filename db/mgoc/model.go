package mgoc

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Model interface {
	Database() string
	Collection() string
	Indexes() ([]mgo.Index, error)
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

func (db *Base) Database() string {
	return db.database
}

func (db *Base) Indexes() ([]mgo.Index, error) {
	return db.Client().Kernel().DB(db.Database()).C(db.Collection()).Indexes()
}

func (db *Base) Collection() string {
	return db.collection
}

func (db *Base) Client() Client {
	return db.client
}

// Do it is used for you to use the nativer mgoc interface according to your own needs,
// Use when you can't find the method you want in this package
func (db *Base) Do(ctx context.Context, f func(c *mgo.Collection) error) error {
	return db.Client().Do(ctx, db, f)
}

func (db *Base) RemoveOne(ctx context.Context, selector interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.Remove(selector)
	})
}
func (db *Base) RemoveId(ctx context.Context, id string) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.RemoveId(bson.ObjectIdHex(id))
	})
}
func (db *Base) RemoveAll(ctx context.Context, selector interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.RemoveAll(selector)
		return err
	})
	return changeInfo, err
}

func (db *Base) DeleteOne(ctx context.Context, selector interface{}) error {
	return db.RemoveOne(ctx, selector)
}

func (db *Base) DeleteId(ctx context.Context, id string) error {
	return db.RemoveId(ctx, id)
}

func (db *Base) DeleteAll(ctx context.Context, selector interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	return db.RemoveAll(ctx, selector)
}

func (db *Base) Insert(ctx context.Context, doc ...interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.Insert(doc...)
	})
}

func (db *Base) UpdateOne(ctx context.Context, selector, update interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.Update(selector, update)
	})
}

func (db *Base) UpdateId(ctx context.Context, id string, update interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(id), update)
	})
}

func (db *Base) UpdateAll(ctx context.Context, selector, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.UpdateAll(selector, update)
		return err
	})
	return
}

func (db *Base) Upsert(ctx context.Context, selector, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.Upsert(selector, update)
		return err
	})
	return
}

func (db *Base) UpsertId(ctx context.Context, id string, update interface{}) (changeInfo *mgo.ChangeInfo, err error) {
	err = db.Do(ctx, func(c *mgo.Collection) error {
		changeInfo, err = c.UpsertId(bson.ObjectIdHex(id), update)
		return err
	})
	return
}

// FindOne the param picker([]string) represents the field to return
func (db *Base) FindOne(ctx context.Context, selector interface{}, picker []string, result interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		query := c.Find(selector)
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(result)
		return err
	})
}

func (db *Base) FindId(ctx context.Context, id string, picker []string, result interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		query := c.FindId(bson.ObjectIdHex(id))
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(result)
		return err
	})
}

func (db *Base) FindAll(ctx context.Context, selector interface{}, sort []string, picker []string, skip int, limit int, result interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
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
		return query.All(result)
	})
}

func (db *Base) FindOneWithSort(ctx context.Context, selector interface{}, sort []string, picker []string, result interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		query := c.Find(selector)
		if sort != nil {
			query = query.Sort(sort...)
		}
		if picker != nil {
			query = query.Select(db.handlePicker(picker))
		}
		err := query.One(result)
		return err
	})
}

func (db *Base) Count(ctx context.Context, selector interface{}) (count int, err error) {
	err = db.Do(ctx, func(c *mgo.Collection) error {
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
func (db *Base) Distinct(ctx context.Context, selector interface{}, key string) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	err := db.Do(ctx, func(c *mgo.Collection) error {
		return c.Find(selector).Distinct(key, &ret)
	})
	return ret, err
}

func (db *Base) PipeAll(ctx context.Context, pipeline interface{}, ret interface{}) error {
	return db.Do(ctx, func(c *mgo.Collection) error {
		return c.Pipe(pipeline).All(ret)
	})
}

func (db *Base) handlePicker(picker []string) interface{} {
	ret := make(bson.M)
	for _, field := range picker {
		ret[field] = 1
	}
	return ret
}
