package ucloud

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"github.com/whereabouts/sdk/oss"
	"io"
)

type client struct {
	config *ufsdk.Config
}

func NewClient(config Config) oss.Client {
	return &client{
		config: &ufsdk.Config{
			PublicKey:       config.PublicKey,
			PrivateKey:      config.PrivateKey,
			BucketName:      config.BucketName,
			BucketHost:      config.BucketHost,
			FileHost:        config.FileHost,
			VerifyUploadMD5: config.VerifyUploadMD5,
		},
	}
}

func (c *client) Upload(file io.Reader, filename string) (string, error) {
	req, err := ufsdk.NewFileRequest(c.config, nil)
	if err != nil {
		return "", err
	}
	err = req.IOPut(file, filename, "")
	if err != nil {
		return string(req.DumpResponse(true)), err
	}
	return req.GetPublicURL(filename), err
}

func (c *client) Download(url string) error {
	req, err := ufsdk.NewFileRequest(c.config, nil)
	if err != nil {
		return err
	}
	// 普通下载 "DownLoadURL"
	err = req.Download(url)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Delete(filename string) error {
	req, err := ufsdk.NewFileRequest(c.config, nil)
	if err != nil {
		return err
	}
	err = req.DeleteFile(filename)
	if err != nil {
		return err
	}
	return nil
}
