package logger

import (
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/sdk/logger/format"
	"github.com/whereabouts/sdk/logger/level"
	"testing"
)

func TestStandardLogger(t *testing.T) {
	logger.Println("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
	logger.New().Println("hello NewLogger")
}

func TestNewLogger(t *testing.T) {
	newLogger := logger.New(
		logger.WithTimestamp(true),
		logger.WithAppName("logger"),
		logger.WithLevel(level.InfoLevelStr),
		logger.WithFormat(format.JSON),
	)
	newLogger.Info("TestNewLogger")
}

func TestResetStandardLogger(t *testing.T) {
	logger.ResetStandardLoggerWithConfig(logger.Config{
		Timestamp: true,
		AppName:   "logger",
		Level:     level.InfoLevelStr,
		Format:    format.JSON,
	})
	logger.ResetStandardLogger(
		logger.WithTimestamp(true),
		logger.WithAppName("logger"),
		logger.WithLevel(level.InfoLevelStr),
		logger.WithFormat(format.JSON),
	)
	logger.Infoln("hello logger")
	logger.Debugln("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
}
