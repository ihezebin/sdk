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
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	// 空间对应机房
	Zone string `json:"zone"`
	// 是否使用https域名
	UseHTTPS bool `json:"use_https"`
	// 上传是否使用CDN上传加速
	UseCdnDomains bool `json:"use_cdn_domains"`
	// 域名地址,包含http://,通过查看外链可以看到,如:http://image-c4lms-qiniu.whereabouts.icu
	Domain string `json:"domain"`
}
