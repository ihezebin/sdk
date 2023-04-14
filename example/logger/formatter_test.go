package logger

import (
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/logger/format"
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
