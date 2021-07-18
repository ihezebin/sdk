package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	StandardLogger().Println("hello logger")
}
