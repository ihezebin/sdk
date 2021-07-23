package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/logger"
	"testing"
)

func TestLevelHook(t *testing.T) {
	logger.StandardLogger().AddHook(logger.NewLocalFSHook(logger.WriterMap{
		logrus.InfoLevel: logger.NewFileWriter("level_hook.log"),
	}, nil))
	logger.Infoln("info")
	logger.Errorln("debug")
}
