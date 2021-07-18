package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/utils/timer"
	"path"
	"path/filepath"
	"runtime"
)

func JSONFormatter() Formatter {
	return convertFormatterReserve(&logrus.JSONFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, filename string) {
			// handle filename
			filename = path.Base(frame.File)
			function = frame.Function
			return
		},
	})
}

func TextFormatter() Formatter {
	return convertFormatterReserve(&logrus.TextFormatter{})
}

func BracketFormatter() Formatter {
	return &bracketFormatter{}
}

type Formatter interface {
	logrus.Formatter
}

func convertFormatter(formatter Formatter) logrus.Formatter {
	return formatter.(logrus.Formatter)
}

func convertFormatterReserve(formatter logrus.Formatter) Formatter {
	return formatter.(Formatter)
}

type bracketFormatter struct{}

func (formatter *bracketFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	time := entry.Time.Format("2006-01-02 15:04:05")
	timestamp := entry.Time.Unix()
	var log string

	if entry.HasCaller() {
		filename := filepath.Base(entry.Caller.File)
		log = fmt.Sprintf("[%s %d] [%s] [%s:%d %s] %s",
			time, timestamp, entry.Level, filename, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		log = fmt.Sprintf("[%s %d] [%s] %s", time, timestamp, entry.Level, entry.Message)
	}

	buffer.WriteString(log)
	return buffer.Bytes(), nil
}
