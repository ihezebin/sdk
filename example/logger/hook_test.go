package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/logger"
	"github.com/whereabouts/sdk-go/logger/hook"
	"github.com/whereabouts/sdk-go/logger/writer"
	"testing"
)

func TestLevelHook(t *testing.T) {
	logger.StandardLogger().AddHook(hook.NewLevelHook(hook.WriterMap{
		logrus.InfoLevel: writer.NewFileWriter("level_hook.log"),
	}, nil))
	logger.Infoln("info")
	logger.Errorln("debug")
}
