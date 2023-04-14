package config

import (
	"fmt"
	"github.com/ihezebin/sdk/utils/stringer"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultTagName = "json"

type ConfigType string

const (
	ConfigTypeJson       ConfigType = "json"
	ConfigTypeToml       ConfigType = "toml"
	ConfigTypeYaml       ConfigType = "yaml"
	ConfigTypeYml        ConfigType = "yml"
	ConfigTypeProperties ConfigType = "properties"
	ConfigTypeProps      ConfigType = "props"
	ConfigTypeProp       ConfigType = "prop"
	ConfigTypeHcl        ConfigType = "hcl"
	ConfigTypeDotenv     ConfigType = "dotenv"
	ConfigTypeEnv        ConfigType = "env"
	ConfigTypeIni        ConfigType = "ini"
)

type Configurator struct {
	kernel     *viper.Viper
	tagName    string
	configType ConfigType
	reader     io.Reader
}

func NewWithFilePath(path string) *Configurator {
	kernel := viper.New()
	kernel.SetConfigFile(handleRelativePath(path))
	return &Configurator{
		kernel:  kernel,
		tagName: defaultTagName,
	}
}

func NewWithReader(reader io.Reader) *Configurator {
	kernel := viper.New()
	return &Configurator{
		kernel:  kernel,
		tagName: defaultTagName,
		reader:  reader,
	}
}

func (c *Configurator) Kernel() *viper.Viper {
	return c.kernel
}

func (c *Configurator) SetTagName(tagName string) *Configurator {
	c.tagName = tagName
	return c
}

func (c *Configurator) SetConfigType(configType ConfigType) *Configurator {
	c.configType = configType
	return c
}

func (c *Configurator) Load(config interface{}) error {
	c.kernel.SetConfigType(string(c.configType))
	if c.kernel.ConfigFileUsed() != "" {
		if err := c.kernel.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to load config file path")
		}
	}
	if c.reader != nil {
		if err := c.kernel.MergeConfig(c.reader); err != nil {
			return errors.Wrap(err, "failed to load config reader")
		}
	}

	return c.Kernel().Unmarshal(config, func(d *mapstructure.DecoderConfig) {
		d.TagName = c.tagName
	})
}

func handleRelativePath(path string) string {
	// handle relative path
	if stringer.NotEmpty(path) && strings.Index(path, ".") == 0 {
		skip := 2
		_, currentPath, _, _ := runtime.Caller(skip)
		fmt.Println(currentPath)
		path = filepath.Join(filepath.Dir(currentPath), path)
	}
	// absolute path
	return path
}
