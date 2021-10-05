package ucloud

type Config struct {
	// "说明1": "管理 bucket 创建和删除必须要公私钥(见 https://console.ucloud.cn/uapi/apikey)，
	// 如果只做文件上传和下载用 TOEKN (见 https://console.ucloud.cn/ufile/token)就够了，为了安全，强烈建议只使用 TOKEN 做文件管理",
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	// "说明2": "以下两个参数是用来管理文件用的。对应的是 file.go 里面的接口，file_host 是不带 bucket 名字的。
	// 比如：北京地域的host填cn-bj.ufileos.com，而不是填 bucketname.cn-bj.ufileos.com。若为自定义域名，请直接带上 http 开头的 URL。
	// 如：http://example.com",
	BucketName string `json:"bucket_name"`
	BucketHost string `json:"bucket_host"`
	FileHost   string `json:"file_host"`
	// "说明3": "verifyUploadMD5 用于数据完整性校验，默认不开启，若要开启请置为true",
	VerifyUploadMD5 bool `json:"verify_upload_md_5"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithPublicKey(publicKey string) Option {
	return func(config *Config) {
		config.PublicKey = publicKey
	}
}

func WithPrivateKey(privateKey string) Option {
	return func(config *Config) {
		config.PrivateKey = privateKey
	}
}

func WithBucketName(bucketName string) Option {
	return func(config *Config) {
		config.BucketName = bucketName
	}
}

func WithBucketHost(bucketHost string) Option {
	return func(config *Config) {
		config.BucketHost = bucketHost
	}
}

func WithFileHost(fileHost string) Option {
	return func(config *Config) {
		config.FileHost = fileHost
	}
}

func WithVerifyUploadMD5(verifyUploadMD5 bool) Option {
	return func(config *Config) {
		config.VerifyUploadMD5 = verifyUploadMD5
	}
}
