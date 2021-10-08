package redisc

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	name string
}

//func (db *Redis) ModelName() string {
//	return db.name
//}

func New(modelName string) *Cache {
	return &Cache{name: modelName}
}

//func (db *Redis) Do(cmd string, args ...interface{}) Result {
//	return db.Client().Do(cmd, args...)
//}

func (db *Cache) Client() Client {
	return getGlobalClient()
}

func (db *Cache) WrapKey(key string) string {
	return fmt.Sprintf("%s_%s", db.name, key)
}

func (db *Cache) Do(cmd string, args ...interface{}) Result {
	return db.Client().Do(cmd, args...)
}

func (db *Cache) Get(key string) Result {
	return db.Client().Do("GET", key)
}

func (db *Cache) Set(key string, val interface{}) Result {
	return db.Client().Do("SET", key, val)
}

func (db *Cache) SetWithExpire(key string, val interface{}, seconds interface{}) Result {
	return db.Client().Do("SETEX", key, seconds, val)
}

func (db *Cache) Incr(key string) Result {
	return db.Client().Do("INCR", key)
}

func (db *Cache) Smembers(key string) Result {
	return db.Client().Do("SMEMBERS", key)
}

func (db *Cache) Hset(key string, hkey string, value interface{}) Result {
	return db.Client().Do("HSET", key, hkey, value)
}

func (db *Cache) HGet(key string, hkey string) Result {
	return db.Client().Do("HGET", key, hkey)
}

func (db *Cache) Delete(key string) Result {
	return db.Client().Do("DEL", key)
}

func (db *Cache) Exists(key string) Result {
	return db.Client().Do("EXISTS", key)
}

func (db *Cache) Expire(key string, seconds interface{}) Result {
	return db.Client().Do("EXPIRE", key, seconds)
}

func (db *Cache) Pipe(method func(c Conn) error) (result Result) {
	conn := Conn{db.Client().GetConn()}
	defer conn.self.Close()
	err := method(conn)
	if err != nil {
		return Result{nil, err}
	}
	err = conn.self.Flush()
	if err != nil {
		return Result{nil, err}
	}
	result.reply, result.err = conn.self.Receive()
	return result
}

type Conn struct {
	self redis.Conn
}

func (conn Conn) Send(command string, args ...interface{}) error {
	return conn.self.Send(command, args...)
}
