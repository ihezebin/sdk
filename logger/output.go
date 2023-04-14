package logger

import (
	"github.com/ihezebin/sdk/logger/writer"
	"io"
	"time"
)

func StandardErrOutput() io.Writer {
	return writer.StandardErrorWriter()
}

func StandardOutput() io.Writer {
	return writer.StandardOutWriter()
}

func FileOutput(filename interface{}) io.Writer {
	return writer.NewFileWriter(filename)
}

func DefaultRotateFileOutput(filename string) io.Writer {
	return writer.DefaultRotateFileWriter(filename)
}

func RotateFileOutput(filename string, rotateTime time.Duration, expireTime time.Duration) io.Writer {
	return writer.NewRotateFileWriter(filename, rotateTime, expireTime)
}
