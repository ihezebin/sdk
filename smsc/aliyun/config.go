package aliyun

type Config struct {
	// 您的AccessKey ID
	AccessKeyId string `mapstructure:"access_key_id" json:"access_key_id"`
	// 您的AccessKey Secret
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret"`
	// 访问的域名
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithAccessKeyId(accessKeyId string) Option {
	return func(config *Config) {
		config.AccessKeyId = accessKeyId
	}
}

func WithAccessKeySecret(accessKeySecret string) Option {
	return func(config *Config) {
		config.AccessKeySecret = accessKeySecret
	}
}

func WithEndpoint(endpoint string) Option {
	return func(config *Config) {
		config.Endpoint = endpoint
	}
}
