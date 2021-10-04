package aliyun

type Config struct {
	// 您的AccessKey ID
	AccessKeyId string `json:"access_key_id"`
	// 您的AccessKey Secret
	AccessKeySecret string `json:"access_key_secret"`
	// 访问的域名
	Endpoint string `json:"endpoint"`
}
