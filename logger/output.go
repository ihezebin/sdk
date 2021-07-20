package logger

import (
	"github.com/whereabouts/sdk-go/logger/writer"
	"os"
)

func StandardErrOutput() *os.File {
	return writer.StandardErrorWriter
}

func StandardOutput() *os.File {
	return writer.StandardOutWriter
}

func FileOutput(filename string) *os.File {
	return writer.NewFileWriter(filename)
}

func RotateFileOutput() *writer.RotateFileWriter {
	return writer.NewRotateFileWriter()
}
