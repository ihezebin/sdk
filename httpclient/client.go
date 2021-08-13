package httpclient

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/httpclient/hook"
	"github.com/whereabouts/sdk/utils/stringer"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

type Client interface {
	Kernel() *resty.Client
	OnBeforeRequest(hooks ...hook.RequestHook) Client
	OnAfterResponse(hooks ...hook.ResponseHook) Client
	NewRequest(ctx context.Context) *resty.Request
	PostJSON(ctx context.Context, path string, values interface{}, headers http.Header, ret interface{}) error
	PostForm(ctx context.Context, path string, values url.Values, headers http.Header, ret interface{}) error
	Get(ctx context.Context, path string, values url.Values, headers http.Header, ret interface{}) error
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

func (c *client) PostJSON(ctx context.Context, path string, values interface{}, headers http.Header, ret interface{}) error {
	r := c.NewRequest(ctx)

	if headers != nil {
		r.Header = headers
	}
	// do something with ctx
	// Todo here
	r.SetHeader("Content-Type", "application/json")

	if values != nil {
		r.SetBody(values)
	}

	if ret != nil {
		r.SetResult(ret)
	}

	resp, err := r.Post(path)

	if err != nil {
		return errors.Errorf("PostJSON:{%s} param:%+v err: %v", r.URL, values, err)
	}
	httpStatusCode := resp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return errors.Errorf("PostJSON:{%s} param:%+v failed: %v", resp.Request.URL, values, httpStatusCode)
	}

	return nil
}

func (c *client) PostForm(ctx context.Context, path string, values url.Values, headers http.Header, ret interface{}) error {
	r := c.NewRequest(ctx)

	if headers != nil {
		r.Header = headers
	}
	// do something with ctx
	// Todo here
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	if values != nil {
		r.FormData = values
	}

	if ret != nil {
		r.SetResult(ret)
	}

	resp, err := r.Post(path)

	if err != nil {
		return errors.Errorf("PostForm:{%s} param:%+v err: %v", r.URL, values, err)
	}
	httpStatusCode := resp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return errors.Errorf("PostForm:{%s} param:%+v failed: %v", resp.Request.URL, values, httpStatusCode)
	}

	return nil
}

func (c *client) Get(ctx context.Context, path string, values url.Values, headers http.Header, ret interface{}) error {
	r := c.NewRequest(ctx)

	if headers != nil {
		r.Header = headers
	}
	// do something with ctx
	// Todo here

	if values != nil {
		r.QueryParam = values
	}

	if ret != nil {
		r.SetResult(ret)
	}

	resp, err := r.Get(path)
	if err != nil {
		return errors.Errorf("Get:{%s} param:%+v err: %v", r.URL, r.QueryParam, err)
	}
	httpStatusCode := resp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return errors.Errorf("Get:{%s} param:%+v failed: %v", resp.Request.URL, values, httpStatusCode)
	}
	return nil
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
