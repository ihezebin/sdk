package logger

import (
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/sdk/logger/format"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	logger.SetFormatter(format.JSONFormatter()).Println("hello JSONFormatter")
}

func TestBracketFormatter(t *testing.T) {
	logger.SetFormatter(format.BracketFormatter()).Println("hello BracketFormatter")
}

func TestTextFormatter(t *testing.T) {
	logger.SetFormatter(format.TextFormatter()).Println("hello TextFormatter")
}
