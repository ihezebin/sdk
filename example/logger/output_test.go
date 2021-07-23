package logger

import (
	"github.com/whereabouts/sdk-go/logger"
	"testing"
	"time"
)

func TestStandardErrOutput(t *testing.T) {
	logger.New().SetOutput(logger.StandardOutput()).Println("TestStandardErrOutput")
}

func TestStandardOutput(t *testing.T) {
	logger.New().SetOutput(logger.StandardErrOutput()).Println("TestStandardOutput")
}

func TestFileOutput(t *testing.T) {
	logger.New().SetOutput(logger.FileOutput("file.log")).Println("TestStandardOutput")
}

func TestRotateFileOutput(t *testing.T) {
	logger := logger.New().SetOutput(
		logger.NewRotateFileWriter("rotate.log").
			SetRotateTime(time.Second * 3).
			SetExpireTime(time.Second * 9))
	for i := 0; i < 12; i++ {
		logger.Println("TestRotateFileOutput")
		time.Sleep(time.Second)
	}
}
