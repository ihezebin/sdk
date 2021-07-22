package writer

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func NewFileWriter(filename string) *os.File {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(errors.Wrap(err, "create log file dir err:"))
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("create file writer failed: %s", err.Error()))
	}
	return file
}
