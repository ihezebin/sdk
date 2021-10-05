package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"io/ioutil"
)

type Client struct {
	kernel *storage.FormUploader
	token  string
	config Config
}

func NewClient(options ...Option) *Client {
	return NewClientWithConfig(newConfig(options...))
}

func NewClientWithConfig(config Config) *Client {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = zoneMap[config.Zone]
	// 是否使用https域名
	cfg.UseHTTPS = config.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = config.UseCdnDomains
	return &Client{
		token:  putPolicy.UploadToken(qbox.NewMac(config.AccessKey, config.SecretKey)),
		kernel: storage.NewFormUploader(&cfg),
		config: config,
	}
}

func (client *Client) Upload(file io.Reader, filename string) (string, error) {
	ret := storage.PutRet{}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = client.kernel.Put(context.Background(), &ret, client.token, filename, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		return "", err
	}
	return storage.MakePublicURL(client.config.Domain, filename), err
}

func (client *Client) Delete(filename string) error {
	mac := qbox.NewMac(client.config.AccessKey, client.config.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: client.config.UseHTTPS,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = zoneMap[client.config.Zone]
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(client.config.Bucket, filename)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) Kernel() *storage.FormUploader {
	return client.kernel
}

func (client *Client) Token() string {
	return client.token
}
