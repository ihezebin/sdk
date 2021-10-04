package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/whereabouts/sdk/ossc"
	"io"
	"io/ioutil"
)

type client struct {
	uploader *storage.FormUploader
	token    string
	config   Config
}

func NewClient(config Config) ossc.Client {
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
	return &client{
		token:    putPolicy.UploadToken(qbox.NewMac(config.AccessKey, config.SecretKey)),
		uploader: storage.NewFormUploader(&cfg),
		config:   config,
	}
}

func (c *client) Upload(file io.Reader, filename string) (string, error) {
	ret := storage.PutRet{}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = c.uploader.Put(context.Background(), &ret, c.token, filename, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		return "", err
	}
	return storage.MakePublicURL(c.config.Domain, filename), err
}

func (c *client) Download(url string) error {
	return nil
}

func (c *client) Delete(filename string) error {
	mac := qbox.NewMac(c.config.AccessKey, c.config.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: c.config.UseHTTPS,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = zoneMap[c.config.Zone]
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(c.config.Bucket, filename)
	if err != nil {
		return err
	}
	return nil
}
