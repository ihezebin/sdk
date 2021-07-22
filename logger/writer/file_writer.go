package writer

import (
	"fmt"
	"os"
)

func NewFileWriter(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(fmt.Sprintf("create file writer failed: %s", err.Error()))
	}
	return file
}
