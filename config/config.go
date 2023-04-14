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

const defaultTagName = ConfigTypeJson
const defaultConfigType = ConfigTypeJson

const (
	ConfigTypeJson       = "json"
	ConfigTypeToml       = "toml"
	ConfigTypeYaml       = "yaml"
	ConfigTypeYml        = "yml"
	ConfigTypeProperties = "properties"
	ConfigTypeProps      = "props"
	ConfigTypeProp       = "prop"
	ConfigTypeHcl        = "hcl"
	ConfigTypeDotenv     = "dotenv"
	ConfigTypeEnv        = "env"
	ConfigTypeIni        = "ini"
)

func New() *Configurator {
	kernel := viper.New()
	kernel.SetConfigType(defaultConfigType)
	return &Configurator{
		kernel:  kernel,
		tagName: defaultTagName,
	}
}

type Configurator struct {
	kernel  *viper.Viper
	tagName string
}

func (c *Configurator) Kernel() *viper.Viper {
	return c.kernel
}

func (c *Configurator) SetTagName(tagName string) *Configurator {
	c.tagName = tagName
	return c
}

func (c *Configurator) SetConfigType(configType string) *Configurator {
	c.Kernel().SetConfigType(configType)
	return c
}

// LoadWithReader need point out the configType, default is json, can use SetConfigType to change
func (c *Configurator) LoadWithReader(reader io.Reader, config interface{}) error {
	if err := c.Kernel().ReadConfig(reader); err != nil {
		return errors.Wrap(err, "read config err")
	}
	if err := c.unmarshal(config); err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	return nil
}

func (c *Configurator) LoadWithFilePath(path string, config interface{}) error {
	c.Kernel().SetConfigFile(c.handleRelativePath(path))
	if err := c.Kernel().ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to load config file path")
	}
	err := c.unmarshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	return nil
}

func (c *Configurator) unmarshal(output interface{}) error {
	return c.Kernel().Unmarshal(output, func(d *mapstructure.DecoderConfig) {
		d.TagName = c.tagName
	})
}

func (c *Configurator) handleRelativePath(path string) string {
	skip := 2
	if c == gConfigurator {
		skip = 3
	}
	if stringer.NotEmpty(path) && strings.Index(path, ".") == 0 {
		_, currentPath, _, _ := runtime.Caller(skip)
		fmt.Println(currentPath)
		path = filepath.Join(filepath.Dir(currentPath), path)
	}
	return path
}
