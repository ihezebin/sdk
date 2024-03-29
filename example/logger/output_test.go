package logger

import (
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/logger/writer"
	"testing"
	"time"
)

func TestStandardErrOutput(t *testing.T) {
	logger.StandardLogger().SetOutput(logger.StandardOutput()).Println("TestStandardErrOutput")
}

func TestStandardOutput(t *testing.T) {
	logger.StandardLogger().SetOutput(logger.StandardErrOutput()).Println("TestStandardOutput")
}

func TestFileOutput(t *testing.T) {
	logger.StandardLogger().SetOutput(logger.FileOutput("file.log")).Println("TestStandardOutput")
}

func TestRotateFileOutput(t *testing.T) {
	logger := logger.StandardLogger().SetOutput(writer.NewRotateFileWriter("rotate.log", time.Second*3, time.Second*9))
	for i := 0; i < 12; i++ {
		logger.Println("TestRotateFileOutput")
		time.Sleep(time.Second)
	}
}
