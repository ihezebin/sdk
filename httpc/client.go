package httpc

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/ihezebin/sdk/httpc/hook"
	"github.com/ihezebin/sdk/utils/stringer"
	"net/http"
	"runtime"
	"time"
)

type Client interface {
	Kernel() *resty.Client
	OnBeforeRequest(hooks ...hook.RequestHook) Client
	OnAfterResponse(hooks ...hook.ResponseHook) Client
	NewRequest(ctx context.Context) *resty.Request
}

type client struct {
	kernel *resty.Client
	config Config
}

func NewClient(options ...Option) (Client, error) {
	return NewClientWithConfig(newConfig(options...))
}

func NewClientWithConfig(config Config) (Client, error) {

	if stringer.NotEmpty(config.Alias) && ClientManager().Has(config.Alias) {
		return nil, errors.Errorf("there is already a client with alias %s, you can choose to use another alias", config.Alias)
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	c := &client{
		kernel: resty.NewWithClient(&http.Client{Transport: transport}),
		config: config,
	}
	c.kernel.SetHostURL(c.config.Host)
	c.kernel.SetTimeout(c.config.Timeout * time.Second)
	c.kernel.SetRetryCount(c.config.RetryCount)
	c.kernel.SetRetryWaitTime(c.config.RetryWaitTime * time.Second)
	c.kernel.SetRetryMaxWaitTime(c.config.RetryMaxWaitTime * time.Second)

	// if alias is not empty, add client to the clientMap
	if stringer.NotEmpty(config.Alias) {
		ClientManager().Add(c.config.Alias, c)
	}

	return c, nil
}

func (c *client) Kernel() *resty.Client {
	return c.kernel
}

func (c *client) OnBeforeRequest(hooks ...hook.RequestHook) Client {
	for _, h := range hooks {
		c.kernel.OnBeforeRequest(resty.RequestMiddleware(h))
	}
	return c
}

func (c *client) OnAfterResponse(hooks ...hook.ResponseHook) Client {
	for _, h := range hooks {
		c.kernel.OnAfterResponse(resty.ResponseMiddleware(h))
	}
	return c
}

func (c *client) NewRequest(ctx context.Context) *resty.Request {
	// do something with ctx
	// Todo here
	return c.kernel.NewRequest().SetContext(ctx)
}

// NewGlobalClient If there is only one http client, you can select the global client
func NewGlobalClient(config Config) (Client, error) {
	var err error
	gClient, err = NewClientWithConfig(config)
	return gClient, err
}

var gClient Client

func GetGlobalClient() Client {
	return gClient
}
