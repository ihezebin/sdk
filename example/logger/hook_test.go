package logger

import (
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/logger/hook"
	"github.com/ihezebin/sdk/logger/level"
	"github.com/ihezebin/sdk/logger/writer"
	"testing"
)

func TestLocalFSHook(t *testing.T) {
	logger.StandardLogger().AddHook(hook.NewLocalFSHook(hook.WriterMap{
		level.InfoLevel: writer.NewFileWriter("localfs_hook.log"),
	}, nil))
	logger.Infoln("TestLocalfsHook")
	logger.Errorln("TestLocalfsHook")
}

func TestMergeHook(t *testing.T) {
	logger.StandardLogger().AddHook(hook.NewMergeHook())
	logger.WithField("name", "korbin").WithFields(field.Fields{
		"age":     18,
		"gender":  "male",
		"address": "cd",
	}).Infoln("TestMergeHook")
	logger.Errorln("TestMergeHook")
}

func TestCallerHook(t *testing.T) {
	callerHook1 := hook.NewCallerHook().SetSimplify(false).SetSource(false)
	logger.StandardLogger().AddHook(callerHook1)
	logger.Infoln("TestCallerHook1")
	logger.Errorln("TestCallerHook1")
	callerHook2 := hook.NewCallerHook().SetSimplify(true).SetSource(true)
	logger.StandardLogger().ReplaceHooks(hook.LevelHooks{
		level.InfoLevel:  []hook.Hook{callerHook2},
		level.ErrorLevel: []hook.Hook{callerHook2},
	})
	logger.Infoln("TestCallerHook2")
	logger.Errorln("TestCallerHook2")
}

func TestFieldsHook(t *testing.T) {
	logger.StandardLogger().AddHook(hook.NewFieldsHook(field.Fields{
		"host":    "127.0.0.1",
		"version": 1.0,
	}))
	logger.Infoln("TestFieldsHook")
	logger.Errorln("TestFieldsHook")
}
