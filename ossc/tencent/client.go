package tencent

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	kernel *cos.Client
	config Config
}

func NewClient(options ...Option) *Client {
	return NewClientWithConfig(newConfig(options...))
}

func NewClientWithConfig(config Config) *Client {
	u, _ := url.Parse(config.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	// 永久密钥
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretID,
			SecretKey: config.SecretKey,
		},
	})
	return &Client{kernel: c, config: config}
}

func (client *Client) Kernel() *cos.Client {
	return client.kernel
}

func (client *Client) Upload(ctx context.Context, file io.Reader, key string) (string, error) {
	_, err := client.kernel.Object.Put(ctx, key, file, nil)
	if err != nil {
		return "", err
	}
	return client.kernel.Object.GetObjectURL(key).String(), err
}

func (client *Client) Delete(ctx context.Context, key string) error {
	_, err := client.kernel.Object.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) DeleteMulti(ctx context.Context, keys ...string) ([]string, error) {
	obs := make([]cos.Object, 0)
	for _, key := range keys {
		obs = append(obs, cos.Object{Key: key})
	}
	opt := &cos.ObjectDeleteMultiOptions{
		Objects: obs,
		// 布尔值，这个值决定了是否启动 Quiet 模式
		// 值为 true 启动 Quiet 模式，值为 false 则启动 Verbose 模式，默认值为 false
		// Quiet: true,
	}

	result, _, err := client.kernel.Object.DeleteMulti(ctx, opt)
	if err != nil {
		return nil, err
	}
	failedKeys := make([]string, 0)
	if result != nil && len(result.Errors) > 0 {
		for _, e := range result.Errors {
			failedKeys = append(failedKeys, e.Key)
		}
	}
	if len(failedKeys) > 0 {
		return failedKeys, errors.New("some key delete failed. detail in failedInfos")
	}

	return nil, nil
}
