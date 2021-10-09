package redisc

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Model interface {
	Name() string
}

type Base struct {
	name   string
	client Client
}

func NewBaseModel(client Client, name string) *Base {
	return &Base{client: client, name: name}
}

func (db *Base) Name() string {
	return db.name
}

func (db *Base) Client() Client {
	return db.client
}

func (db *Base) Key(key string) string {
	return fmt.Sprintf("%s_%s", db.Name(), key)
}

func (db *Base) Do(cmd string, args ...interface{}) Result {
	return db.Client().Do(cmd, args...)
}

func (db *Base) Get(key string) Result {
	return db.Client().Do("GET", key)
}

func (db *Base) Set(key string, val interface{}) Result {
	return db.Client().Do("SET", key, val)
}

func (db *Base) SetIfNotExists(key string, val interface{}) bool {
	return db.Client().Do("SETNX", key, val).Bool()
}

func (db *Base) SetWithExpire(key string, val interface{}, seconds interface{}) Result {
	return db.Client().Do("SETEX", key, seconds, val)
}

func (db *Base) IncrBy(key string, num int) int {
	return db.Client().Do("INCRBY", key, num).Int()
}

func (db *Base) DecrBy(key string, num int) int {
	return db.Client().Do("DECRBY", key, num).Int()
}

func (db *Base) LPush(key string, val interface{}) Result {
	return db.Client().Do("LPUSH", key, val)
}

func (db *Base) RPush(key string, val interface{}) Result {
	return db.Client().Do("RPUSH", key, val)
}

func (db *Base) LPop(key string) Result {
	return db.Client().Do("LPOP", key)
}

func (db *Base) RPop(key string) Result {
	return db.Client().Do("RPOP", key)
}

func (db *Base) LRange(key string, start int, end int) Result {
	return db.Client().Do("LRANGE", key, start, end)
}

func (db *Base) SMembers(key string) Result {
	return db.Client().Do("SMEMBERS", key)
}

func (db *Base) HSet(key string, hKey string, value interface{}) Result {
	return db.Client().Do("HSET", key, hKey, value)
}

func (db *Base) HGet(key string, hKey string) Result {
	return db.Client().Do("HGET", key, hKey)
}

func (db *Base) Delete(key string) Result {
	return db.Client().Do("DEL", key)
}

func (db *Base) Exists(key string) Result {
	return db.Client().Do("EXISTS", key)
}

func (db *Base) Expire(key string, seconds interface{}) Result {
	return db.Client().Do("EXPIRE", key, seconds)
}

func (db *Base) Pipe(method func(c Conn) error) (result Result) {
	conn := Conn{db.Client().GetConn()}
	defer conn.kernel.Close()
	err := method(conn)
	if err != nil {
		return Result{nil, err}
	}
	err = conn.kernel.Flush()
	if err != nil {
		return Result{nil, err}
	}
	result.reply, result.err = conn.kernel.Receive()
	return result
}

type Conn struct {
	kernel redis.Conn
}

func (conn Conn) Send(command string, args ...interface{}) error {
	return conn.kernel.Send(command, args...)
}
