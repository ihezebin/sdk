package mapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk-go/utils/timer"
	"reflect"
)

func Json2Map(s string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &m)
	return m, err
}

func Map2Json(m interface{}) (string, error) {
	v := reflect.ValueOf(m)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Map {
		return "", errors.New("the param is not a map")
	}
	marshal, err := json.Marshal(v.Interface())
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

func Map2Struct(m interface{}, v interface{}) error {
	s, err := Map2Json(m)
	if err != nil {
		return err
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("the v(%+v) is not a pointer", v))
	}
	err = json.Unmarshal([]byte(s), v)
	return err
}

func Struct2Map(v interface{}) (map[string]interface{}, error) {
	value := reflect.ValueOf(v)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, errors.New("the param is not a struct")
	}
	t := reflect.TypeOf(value.Interface())
	var m = make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		key := t.Field(i).Tag.Get("bson")
		if key == "" {
			key = t.Field(i).Tag.Get("json")
			if key == "" {
				key = t.Field(i).Name
			}
		}
		m[key] = value.Field(i).Interface()
	}
	return m, nil
}

func IsKeyValid(m bson.M, key string) bool {
	_, ok := m[key]
	return ok
}

func IsKeyTimeValid(m bson.M, key string) bool {
	value, ok := m[key]
	if !ok {
		return false
	}
	if !timer.IsStrTime(value) && !timer.IsTime(value) && !timer.IsUnixTime(value) {
		return false
	}
	return true
}

func BsonM2Map(m bson.M) map[string]interface{} {
	return m
}

func Map2BsonM(m map[string]interface{}) bson.M {
	return m
}

func Struct2BsonM(v interface{}) (bson.M, error) {
	m, err := Struct2Map(v)
	if err != nil {
		return nil, err
	}
	return Map2BsonM(m), nil
}

func BsonM2Struct(m bson.M, v interface{}) error {
	return Map2Struct(BsonM2Map(m), v)
}
