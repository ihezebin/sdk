package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"io/ioutil"
)

var (
	ZoneMap = map[string]*storage.Region{
		"huanan":  &storage.ZoneHuanan,
		"huadong": &storage.ZoneHuadong,
		"huabei":  &storage.ZoneHuabei,
	}
)

const (
	ZoneHuanan  = "huanan"
	ZoneHuadong = "huadong"
	ZoneHuabei  = "huabei"
)

type Config struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	// 空间对应机房
	Zone string `json:"zone"`
	// 是否使用https域名
	UseHTTPS bool `json:"use_https"`
	// 上传是否使用CDN上传加速
	UseCdnDomains bool `json:"use_cdn_domains"`
	// 域名地址,包含http://,通过查看外链可以看到,如:http://image-c4lms-qiniu.whereabouts.icu
	Domain string `json:"domain"`
}

type Client struct {
	Uploader *storage.FormUploader
	Token    string
	Config   *Config
}

// 如果存在同名且相同的文件,则直接返回url
func (cloud *Client) Upload(file io.Reader, fileName string) (string, error) {
	ret := storage.PutRet{}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = cloud.Uploader.Put(context.Background(), &ret, cloud.Token, fileName, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		return "", err
	}
	return storage.MakePublicURL(cloud.Config.Domain, fileName), err
}

func (cloud *Client) Download(fileName string) error {
	return nil
}

func (cloud *Client) Delete(fileName string) error {
	mac := qbox.NewMac(cloud.Config.AccessKey, cloud.Config.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: cloud.Config.UseHTTPS,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = ZoneMap[cloud.Config.Zone]
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(cloud.Config.Bucket, fileName)
	if err != nil {
		return err
	}
	return nil
}
