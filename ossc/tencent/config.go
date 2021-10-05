package tencent

type Config struct {
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	BucketURL string `json:"bucket_url"`
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
