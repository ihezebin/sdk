package configure

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const DefaultConfigPath = "./config/application.json"

var callerSkip = 1

func LoadJSON(path string, conf interface{}) error {
	_, currentPath, _, _ := runtime.Caller(callerSkip)
	path = filepath.Join(filepath.Dir(currentPath), path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Errorf("configure file %s not exist", path)
	}
	fmt.Println(path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return err
	}
	return nil
}

func LoadJSONWithCmd(conf interface{}) error {
	callerSkip = 2
	path := GetCmdParam("c", DefaultConfigPath)
	err := LoadJSON(path.String(), conf)
	if err != nil {
		return err
	}
	callerSkip = 1
	return nil
}
