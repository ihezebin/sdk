package ucloud

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"io"
)

type Config struct {
	// "说明1": "管理 bucket 创建和删除必须要公私钥(见 https://console.ucloud.cn/uapi/apikey)，
	// 如果只做文件上传和下载用 TOEKN (见 https://console.ucloud.cn/ufile/token)就够了，为了安全，强烈建议只使用 TOKEN 做文件管理",
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	// "说明2": "以下两个参数是用来管理文件用的。对应的是 file_writer.go 里面的接口，file_host 是不带 bucket 名字的。
	// 比如：北京地域的host填cn-bj.ufileos.com，而不是填 bucketname.cn-bj.ufileos.com。若为自定义域名，请直接带上 http 开头的 URL。
	// 如：http://example.com",
	BucketName string `json:"bucket_name"`
	BucketHost string `json:"bucket_host"`
	FileHost   string `json:"file_host"`
	// "说明3": "verifyUploadMD5 用于数据完整性校验，默认不开启，若要开启请置为true",
	VerifyUploadMD5 bool `json:"verify_upload_md_5"`
}

type Client struct {
	Config *ufsdk.Config
}

// 如果存在同名且相同的文件,则直接返回url
func (cloud *Client) Upload(file io.Reader, fileName string) (string, error) {
	req, err := ufsdk.NewFileRequest(cloud.Config, nil)
	if err != nil {
		return "", err
	}
	err = req.IOPut(file, fileName, "")
	if err != nil {
		return string(req.DumpResponse(true)), err
	}
	return req.GetPublicURL(fileName), err
}

func (cloud *Client) Download(url string) error {
	req, err := ufsdk.NewFileRequest(cloud.Config, nil)
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

func (cloud *Client) Delete(fileName string) error {
	req, err := ufsdk.NewFileRequest(cloud.Config, nil)
	if err != nil {
		return err
	}
	err = req.DeleteFile(fileName)
	if err != nil {
		return err
	}
	return nil
}
