package redis

import (
	"github.com/gomodule/redigo/redis"
)

type Client interface {
	GetConn() (conn redis.Conn)
	Close()
	Do(cmd string, args ...interface{}) Result
}

var globalClient Client

func getGlobalClient() Client {
	return globalClient
}

func Init(config Config) (Client, error) {
	c, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	globalClient = c
	return c, err
}

func NewClient(config Config) (Client, error) {
	dialOption := make([]redis.DialOption, 0)
	if config.Username != "" {
		dialOption = append(dialOption, redis.DialUsername(config.Username))
	}
	if config.Password != "" {
		dialOption = append(dialOption, redis.DialPassword(config.Password))
	}
	if config.ClientName != "" {
		dialOption = append(dialOption, redis.DialClientName(config.ClientName))
	}
	if config.Database != 0 {
		dialOption = append(dialOption, redis.DialDatabase(config.Database))
	}
	var err error
	c := &client{
		config: config,
		pool: &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: config.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				conn, er := redis.Dial(defaultNetwork, config.Addr, dialOption...)
				err = er
				return conn, er
			},
		},
	}
	globalClient = c
	return c, err
}

type client struct {
	pool   *redis.Pool
	config Config
}

func (c *client) GetConn() (conn redis.Conn) {
	return c.pool.Get()
}

func (c *client) Close() {
	c.pool.Close()
}

func (c *client) Do(cmd string, args ...interface{}) Result {
	conn := c.GetConn()
	defer conn.Close()
	reply, err := conn.Do(cmd, args...)
	return Result{reply, err}
}
