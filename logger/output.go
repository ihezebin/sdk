package logger

import (
	"io"
	"time"
)

func StandardErrOutput() io.Writer {
	return StandardErrorWriter()
}

func StandardOutput() io.Writer {
	return StandardOutWriter()
}

func FileOutput(filename interface{}) io.Writer {
	return NewFileWriter(filename)
}

func DefaultRotateFileOutput(filename string) io.Writer {
	return DefaultRotateFileWriter(filename)
}

func RotateFileOutput(filename string, rotateTime time.Duration, expireTime time.Duration) io.Writer {
	return NewRotateFileWriter(filename, rotateTime, expireTime)
}
