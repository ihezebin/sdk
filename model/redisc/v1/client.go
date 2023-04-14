package v1

import (
	"github.com/gomodule/redigo/redis"
)

type Client interface {
	GetConn() (conn redis.Conn)
	Close()
	Do(cmd string, args ...interface{}) Result
}

func NewClient(options ...Option) (Client, error) {
	return NewClientWithConfig(newConfig(options...))
}

func NewClientWithConfig(config Config) (Client, error) {
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
	c := &client{
		config: config,
		pool: &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: config.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.Dial(defaultNetwork, config.Addr, dialOption...)
			},
		},
	}
	// 测试是否连接成功
	conn, err := c.pool.Dial()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	return c, nil
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
