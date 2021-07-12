package test

import (
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk-go/model/mongoc"
)

var (
	userDB = &UserDB{mongoc.New("test", "user", User{})}
)

func GetUserDB() *UserDB {
	return userDB
}

type UserDB struct {
	*mongoc.MongoDB
}

type User struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
	Age  int           `json:"age"`
}

func (user *UserDB) Add(users ...interface{}) error {
	return user.Insert(users...)
}

func (user *UserDB) AddOne(u User) error {
	return user.Insert(u)
}

func (user *UserDB) GetByName(name string) (*User, error) {
	u := &User{}
	selector := bson.M{"name": name}
	err := user.FindOne(selector, nil, u)
	if err != nil {
		return nil, err
	}
	return u, err
}

func (user *UserDB) GetAll() ([]*User, error) {
	ret := make([]*User, 0)
	err := user.FindAll(nil, nil, nil, 0, 0, &ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func (user *UserDB) ModifyAgeByName(name string, age int) error {
	u := &User{}
	selector := bson.M{"name": name}
	return user.Modify(selector, bson.M{"age": age}, u)
}

func (user *UserDB) DeleteByAge(age int) error {
	selector := bson.M{"age": age}
	return user.Remove(selector)
}

func (user *UserDB) DeleteTest(age int) error {
	selector := bson.M{"age": age}
	return user.Modify(selector, nil, nil, true)
}

func (user *UserDB) ReplaceByName(name string, u User) error {
	return user.Replace(bson.M{"name": name}, u)
}
