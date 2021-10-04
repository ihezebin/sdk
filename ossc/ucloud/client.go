package ucloud

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"io"
)

type Client interface {
	Upload(file io.Reader, key string) (string, error)
	Delete(key string) error
}

type client struct {
	config *ufsdk.Config
}

func NewClient(config Config) Client {
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

func (c *client) Upload(file io.Reader, key string) (string, error) {
	req, err := ufsdk.NewFileRequest(c.config, nil)
	if err != nil {
		return "", err
	}
	err = req.IOPut(file, key, "")
	if err != nil {
		return string(req.DumpResponse(true)), err
	}
	return req.GetPublicURL(key), err
}

func (c *client) Delete(key string) error {
	req, err := ufsdk.NewFileRequest(c.config, nil)
	if err != nil {
		return err
	}
	err = req.DeleteFile(key)
	if err != nil {
		return err
	}
	return nil
}
