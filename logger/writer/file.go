package writer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

// NewFileWriter file can be a string, io.Writer
func NewFileWriter(file interface{}) io.Writer {
	switch file.(type) {
	case string:
		filename := file.(string)
		dir := filepath.Dir(filename)
		_ = os.MkdirAll(dir, os.ModePerm)
		writer, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(fmt.Sprintf("create file writer failed: %s", err.Error()))
		}
		return writer
	case io.Writer:
		return file.(io.Writer)
	default:
		panic(fmt.Sprintf("unsupported file type: %v", reflect.TypeOf(file)))
	}
}
