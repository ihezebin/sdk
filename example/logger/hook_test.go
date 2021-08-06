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
	logger.Infoln("TestLocalfsHook")
	logger.Errorln("TestLocalfsHook")
}

func TestMergeHook(t *testing.T) {
	logger.StandardLogger().AddHook(logger.NewMergeHook())
	logger.WithField("name", "korbin").WithFields(logger.Fields{
		"age":     18,
		"gender":  "male",
		"address": "cd",
	}).Infoln("TestMergeHook")
	logger.Errorln("TestMergeHook")
}

func TestCallerHook(t *testing.T) {
	callerHook1 := logger.NewCallerHook().SetSimplify(false).SetSource(false)
	logger.StandardLogger().AddHook(callerHook1)
	logger.Infoln("TestCallerHook1")
	logger.Errorln("TestCallerHook1")
	callerHook2 := logger.NewCallerHook().SetSimplify(true).SetSource(true)
	logger.StandardLogger().ReplaceHooks(logger.LevelHooks{
		logger.InfoLevel:  []logger.Hook{callerHook2},
		logger.ErrorLevel: []logger.Hook{callerHook2},
	})
	logger.Infoln("TestCallerHook2")
	logger.Errorln("TestCallerHook2")
}

func TestFieldsHook(t *testing.T) {
	logger.StandardLogger().AddHook(logger.NewFieldsHook(logger.Fields{
		"host":    "127.0.0.1",
		"version": 1.0,
	}))
	logger.Infoln("TestFieldsHook")
	logger.Errorln("TestFieldsHook")
}
