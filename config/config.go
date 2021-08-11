package config

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/whereabouts/sdk/cli"
	"github.com/whereabouts/sdk/utils/stringer"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

var defaultConfigPath = "./config.json"

func New() *Configurator {
	return &Configurator{viper.New()}
}

var gConfigurator = New()

func Load(file interface{}, config interface{}) error {
	return gConfigurator.Load(file, config)
}

func LoadWithDefault(config interface{}) error {
	return gConfigurator.LoadWithDefault(config)
}

func LoadWithCli(key, value string, config interface{}) error {
	return gConfigurator.LoadWithCli(key, value, config)
}

func LoadWithEnv(key string, config interface{}) error {
	return gConfigurator.LoadWithEnv(key, config)
}

//func LoadWithURI(uri string, config interface{}) error {
//	return gConfigurator.LoadWithURI(uri, config)
//}

func SetDefaultConfigPath(path string) {
	defaultConfigPath = path
}

type Configurator struct {
	kernel *viper.Viper
}

func (c *Configurator) Kernel() *viper.Viper {
	return c.kernel
}

func (c *Configurator) Load(file interface{}, config interface{}) error {
	switch file.(type) {
	case string:
		return c.loadFilePath(file.(string), config)
	case io.Reader:
		return c.loadFileReader(file.(io.Reader), config)
	case []byte:
		return c.loadFileReader(bytes.NewReader(file.([]byte)), config)
	default:
		return errors.Errorf("unsupported config file type: %v", reflect.TypeOf(file))
	}
}

func (c *Configurator) LoadWithDefault(config interface{}) error {
	return c.loadFilePath(defaultConfigPath, config)
}

func (c *Configurator) LoadWithCli(key, value string, config interface{}) error {
	cli.WithFlagString(key, value, "config file path")
	cliV, err := cli.Run()
	if err != nil {
		return err
	}
	return c.loadFilePath(cliV.String(key), config)
}

func (c *Configurator) LoadWithEnv(key string, config interface{}) error {
	return c.loadFilePath(os.Getenv(key), config)
}

//func (c *Configurator) LoadWithURI(path string, config interface{}) error {
//	return nil
//}

func (c *Configurator) loadFileReader(reader io.Reader, config interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read config reader data")
	}
	for _, configType := range viper.SupportedExts {
		c.Kernel().SetConfigType(configType)
		if err = c.Kernel().ReadConfig(bytes.NewReader(data)); err != nil {
			continue
		} else {
			err = c.Kernel().Unmarshal(config)
			if err != nil {
				return errors.Wrap(err, "failed to unmarshal config")
			}
			// check if zero value
			configV := reflect.ValueOf(config)
			for configV.Kind() == reflect.Ptr {
				configV = configV.Elem()
			}
			if configV.IsZero() {
				continue
			}
			return nil
		}
	}
	return errors.New("failed to load config reader, the data format does not match these: [json, toml, yaml, yml, properties, props, prop, hcl, dotenv, env, ini]")
}

func (c *Configurator) loadFilePath(path string, config interface{}) error {
	skip := 3
	if c == gConfigurator {
		skip = 4
	}
	c.Kernel().SetConfigFile(c.handleRelativePath(path, skip))
	if err := c.Kernel().ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to load config file path")
	}
	err := c.Kernel().Unmarshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	return nil
}

func (c *Configurator) handleRelativePath(path string, skip int) string {
	if stringer.NotEmpty(path) && stringer.Equals(path[:1], ".") {
		_, currentPath, _, _ := runtime.Caller(skip)
		path = filepath.Join(filepath.Dir(currentPath), path)
	}
	return path
}
