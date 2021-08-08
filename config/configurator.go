package config

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var defaultConfigPath = "./config/config.json"

type Configurator struct {
	kernel *viper.Viper
}

func New() *Configurator {
	return &Configurator{viper.New()}
}

func (c *Configurator) Kernel() *viper.Viper {
	return c.kernel
}

func (c *Configurator) Load(file interface{}, config interface{}) error {
	switch file.(type) {
	case string:
		c.Kernel().SetConfigFile(file.(string))
		if err := c.Kernel().ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read config file")
		}
	case *os.File:
		c.Kernel().SetConfigType(strings.Trim(filepath.Ext(file.(*os.File).Name()), "."))
		if err := c.Kernel().ReadConfig(file.(io.Reader)); err != nil {
			return errors.Wrap(err, "failed to read config reader")
		}
	case io.Reader:
		//c.Kernel().SetConfigType("")
		if err := c.Kernel().ReadConfig(file.(io.Reader)); err != nil {
			return errors.Wrap(err, "failed to read config reader")
		}
	default:
		return errors.Errorf("unsupported file type: %v", reflect.TypeOf(file))
	}

	err := c.Kernel().Unmarshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	return nil
}

func (c *Configurator) LoadWithDefault(config interface{}) error {
	return c.Load(defaultConfigPath, config)
}

func (c *Configurator) LoadWithCmd(key, value string, config interface{}) error {
	file := flag.String(key, value, "go run main.go -[name] [value]")
	flag.Parse()
	return c.Load(*file, config)
}

func (c *Configurator) LoadWithEnv(key string, config interface{}) error {
	return c.Load(os.Getenv(key), config)
}

//func (c *Configurator) LoadWithURI(path string, config interface{}) error {
//	return nil
//}
