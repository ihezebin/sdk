package redisc

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type wrapClient = redis.UniversalClient

type Client struct {
	wrapClient
}

func NewClient(ctx context.Context, config *Config) (*Client, error) {
	options := config.Convert2Otions()
	client := redis.NewUniversalClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{wrapClient: client}, nil
}

func NewClientWithOptions(ctx context.Context, options ...Option) (*Client, error) {
	config := newConfig(options...)
	return NewClient(ctx, config)
}

var gClient *Client

// NewGlobalClient If there is only one Mongo, you can select the global client
func NewGlobalClient(ctx context.Context, config *Config) (*Client, error) {
	var err error
	if gClient, err = NewClient(ctx, config); err != nil {
		return nil, err
	}
	return gClient, nil
}

func GetGlobalClient() *Client {
	return gClient
}

func (c *Client) Kernel() redis.UniversalClient {
	return c.wrapClient
}
