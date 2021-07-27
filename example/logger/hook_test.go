package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

func TestLocalfsHook(t *testing.T) {
	logger.StandardLogger().AddHook(logger.NewLocalFSHook(logger.WriterMap{
		logrus.InfoLevel: logger.NewFileWriter("localfs_hook.log"),
	}, nil))
	logger.Infoln("info")
	logger.Errorln("debug")
}
