package logger

import (
	"github.com/whereabouts/sdk-go/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.Println("hello logger")
	logger.StandardLogger().Println("hello standard")
	logger.New().Println("hello new")
}
