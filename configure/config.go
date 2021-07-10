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

func LoadJSON(path string, conf interface{}) error {
	_, currentPath, _, _ := runtime.Caller(1)
	path = filepath.Join(filepath.Dir(currentPath), path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("config file not exist")
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
