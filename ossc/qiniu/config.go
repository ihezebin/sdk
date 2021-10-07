package qiniu

import "github.com/qiniu/go-sdk/v7/storage"

var (
	zoneMap = map[string]*storage.Region{
		ZoneHuanan:  &storage.ZoneHuanan,
		ZoneHuadong: &storage.ZoneHuadong,
		ZoneHuabei:  &storage.ZoneHuabei,
	}
)

const (
	ZoneHuanan  = "huanan"
	ZoneHuadong = "huadong"
	ZoneHuabei  = "huabei"
)

type Config struct {
	AccessKey string `mapstructure:"access_key" json:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
	// 空间对应机房
	Zone string `mapstructure:"zone" json:"zone"`
	// 是否使用https域名
	UseHTTPS bool `mapstructure:"use_https" json:"use_https"`
	// 上传是否使用CDN上传加速
	UseCdnDomains bool `mapstructure:"use_cdn_domains" json:"use_cdn_domains"`
	// 域名地址,包含http://,通过查看外链可以看到,如:http://image-c4lms-qiniu.whereabouts.icu
	Domain string `mapstructure:"domain" json:"domain"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithAccessKey(accessKey string) Option {
	return func(config *Config) {
		config.AccessKey = accessKey
	}
}

func WithSecretKey(secretKey string) Option {
	return func(config *Config) {
		config.SecretKey = secretKey
	}
}

func WithBucket(bucket string) Option {
	return func(config *Config) {
		config.Bucket = bucket
	}
}

func WithZone(zone string) Option {
	return func(config *Config) {
		config.Zone = zone
	}
}

func WithUseHTTPS(useHTTPS bool) Option {
	return func(config *Config) {
		config.UseHTTPS = useHTTPS
	}
}

func WithUseCdnDomains(useCdnDomains bool) Option {
	return func(config *Config) {
		config.UseCdnDomains = useCdnDomains
	}
}

func WithDomain(domain string) Option {
	return func(config *Config) {
		config.Domain = domain
	}
}
