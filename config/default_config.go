package configure

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/chassis/logger"
	"io/ioutil"
	"reflect"
)

const (
	fieldNamePort     = "Port"
	fieldNameEnv      = "Env"
	defaultConfigPath = "./config/config.json"
	//defaultInternalConfigPath = "./engine/configure/config.json"
)

func init() {
	//data, err := ioutil.ReadFile(defaultInternalConfigPath)
	//if err != nil {
	//	logger.Fatalf("read internal default config file err: %+v", err)
	//}
	//err = json.Unmarshal(data, &dConfig)
	//if err != nil {
	//	logger.Fatalf("unmarshal internal default config err: %+v", err)
	//}
	dConfig.Env = gin.DebugMode
	dConfig.Port = 8080
}

type IConfig interface {
	Set()
}

type DefaultConfig struct {
	Port int64  `json:"port"`
	Env  string `json:"env"`
}

var dConfig = DefaultConfig{}

func Load(conf IConfig) (*DefaultConfig, error) {
	// first, find the path from the command line parameter or the default directory
	cmdParam := GetCmdParam("c", defaultConfigPath)
	data, err := ioutil.ReadFile(cmdParam.String())
	// if fail to read the default configure file, that will appoint some default values
	if err != nil {
		if cmdParam.String() == defaultConfigPath {
			logger.Warning("read file: ./config/config.json err, but the default configuration is selected!")
			return DefaultLoad(), nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	confV := reflect.ValueOf(conf).Elem()
	if port := confV.FieldByName(fieldNamePort); port.IsValid() && port.Int() > reflect.ValueOf(0).Int() {
		dConfig.Port = port.Int()
	}
	if env := confV.FieldByName(fieldNameEnv); env.IsValid() && env.String() != reflect.ValueOf("").String() {
		dConfig.Env = env.String()
	}
	return &dConfig, nil
}

func DefaultLoad() *DefaultConfig {
	return &dConfig
}
