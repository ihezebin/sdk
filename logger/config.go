package logger

import (
	"time"
)

type Config struct {
	// Level default INFO
	Level string `json:"level"`

	// AppName Application name or web service name. If it is not empty, a field named "app" will be added to the log
	// 应用名称或Web服务名称，非空时会在日志中添加名为"app"的Field
	AppName string `json:"app_name"`

	// Format The default format is JSON. Formatters.go provides two additional formats: text and bracket
	// 默认格式化方式为JSON，formatters.go额外提供了"text"和"bracket"两种格式
	Format string `json:"format"`

	// Timestamp Whether to display the timestamp, default true
	// 是否显示时间戳, 默认true
	Timestamp bool `json:"timestamp"`

	// File If this item is configured, it means that the output is a local file
	// 配置了File表示output为本地文件
	File string `json:"file"`

	// ErrFile If this item is configured, the log at error level will be input to a separate local file
	// 配置了ErrFile表示error级别的日志将输入到单独的本地文件
	ErrFile string `json:"err_file"`

	// RotateFile If this item is configured, which means that output is divided to the local file system by rotation
	// If File is configured at the same time, RotateFile will prevail
	// 配置了RotateFile表示output采用轮转分割到本地文件系统，若同时配置了File, 将以RotateFile为准
	RotateFile RotateFileConfig `json:"rotate_file"`
}

type RotateFileConfig struct {
	File string `json:"file"`
	// RotateTime Log cutting interval
	// 日志切割时间间隔
	RotateTime time.Duration `json:"rotate_time"`
	// ExpireTime Maximum storage time of log files
	// 日志文件最大保存时间
	ExpireTime time.Duration `json:"expire_time"`
}

type ConfigFunc func(config *Config)

func newConfig(cfs ...ConfigFunc) Config {
	config := Config{}
	for _, conf := range cfs {
		conf(&config)
	}
	return config
}

func WithAppName(appName string) ConfigFunc {
	return func(config *Config) {
		config.AppName = appName
	}
}

func WithLevel(level string) ConfigFunc {
	return func(config *Config) {
		config.Level = level
	}
}

func WithFormat(format string) ConfigFunc {
	return func(config *Config) {
		config.Format = format
	}
}

func WithTimestamp(timestamp bool) ConfigFunc {
	return func(config *Config) {
		config.Timestamp = timestamp
	}
}

func WithFile(file string) ConfigFunc {
	return func(config *Config) {
		config.File = file
	}
}

func WithErrFile(errFile string) ConfigFunc {
	return func(config *Config) {
		config.ErrFile = errFile
	}
}

func WithRotateFile(rotateFile RotateFileConfig) ConfigFunc {
	return func(config *Config) {
		config.RotateFile = rotateFile
	}
}