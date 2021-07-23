package logger

import (
	"io"
)

func StandardErrOutput() io.Writer {
	return StandardErrorWriter
}

func StandardOutput() io.Writer {
	return StandardOutWriter
}

func FileOutput(filename string) io.Writer {
	return NewFileWriter(filename)
}

func RotateFileOutput(filename string) io.Writer {
	return NewRotateFileWriter(filename)
}
