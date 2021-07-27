package oss

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"github.com/whereabouts/sdk/oss/qiniu"
	"github.com/whereabouts/sdk/oss/ucloud"
	"io"
)

type Client interface {
	Upload(file io.Reader, fileName string) (string, error)
	Download(url string) error
	Delete(fileName string) error
}

func NewUCloudClient(config *ucloud.Config) Client {
	return &ucloud.Client{
		Config: &ufsdk.Config{
			PublicKey:       config.PublicKey,
			PrivateKey:      config.PrivateKey,
			BucketName:      config.BucketName,
			BucketHost:      config.BucketHost,
			FileHost:        config.FileHost,
			VerifyUploadMD5: config.VerifyUploadMD5,
		},
	}
}

func NewQiNiuClient(config *qiniu.Config) Client {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = qiniu.ZoneMap[config.Zone]
	// 是否使用https域名
	cfg.UseHTTPS = config.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = config.UseCdnDomains
	return &qiniu.Client{
		Token:    putPolicy.UploadToken(qbox.NewMac(config.AccessKey, config.SecretKey)),
		Uploader: storage.NewFormUploader(&cfg),
		Config:   config,
	}
}
