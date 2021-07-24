package logger

import (
	"github.com/whereabouts/sdk-go/logger"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	logger.SetFormatter(logger.JSONFormatter()).Println("hello WebFormatter")
}

func TestBracketFormatter(t *testing.T) {
	logger.SetFormatter(logger.BracketFormatter()).Println("hello BracketFormatter")
}

func TestTextFormatter(t *testing.T) {
	logger.SetFormatter(logger.TextFormatter()).Println("hello TextFormatter")
}
