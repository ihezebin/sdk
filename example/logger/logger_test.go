package logger

import (
	"github.com/whereabouts/sdk/logger"
	"testing"
)

func TestStandardLogger(t *testing.T) {
	logger.Println("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
	logger.NewLogger().Println("hello NewLogger")
}

func TestNewLogger(t *testing.T) {
	newLogger := logger.NewLogger(
		logger.WithTimestamp(true),
		logger.WithAppName("logger"),
		logger.WithLevel(logger.InfoLevelStr),
		logger.WithFormat(logger.FormatJSON),
	)
	newLogger.Info("TestNewLogger")
}

func TestResetStandardLogger(t *testing.T) {
	logger.ResetStandardLoggerWithConfig(logger.Config{
		Timestamp: true,
		AppName:   "logger",
		Level:     logger.InfoLevelStr,
		Format:    logger.FormatJSON,
	})
	logger.ResetStandardLogger(
		logger.WithTimestamp(true),
		logger.WithAppName("logger"),
		logger.WithLevel(logger.InfoLevelStr),
		logger.WithFormat(logger.FormatJSON),
	)
	logger.Println("hello logger")
	logger.Debugln("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
}
