package tencent

type Config struct {
	SecretID  string `mapstructure:"secret_id" json:"secret_id"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	BucketURL string `mapstructure:"bucket_url" json:"bucket_url"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithSecretID(secretID string) Option {
	return func(config *Config) {
		config.SecretID = secretID
	}
}

func WithSecretKey(secretKey string) Option {
	return func(config *Config) {
		config.SecretKey = secretKey
	}
}

func WithBucketURL(bucketURL string) Option {
	return func(config *Config) {
		config.BucketURL = bucketURL
	}
}
