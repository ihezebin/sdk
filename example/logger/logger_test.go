package logger

import (
	"github.com/whereabouts/sdk-go/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.Println("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
	logger.NewLogger().Println("hello NewLogger")
}

func TestResetStandardLogger(t *testing.T) {
	logger.ResetStandardLoggerWithConfig(logger.Config{
		Timestamp: true,
		AppName:   "logger",
		Level:     logger.InfoLevelString,
		Format:    logger.FormatJSON,
	})
	logger.Println("hello logger")
	logger.Debugln("hello logger")
	logger.StandardLogger().Println("hello StandardLogger")
}
