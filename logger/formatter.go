package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk/utils/timer"
	"strings"
)

type FormatType = string

const (
	FormatJSON    FormatType = "json"
	FormatText    FormatType = "text"
	FormatBracket FormatType = "bracket"
)

var (
	formatterMap = map[FormatType]Formatter{
		FormatJSON:    JSONFormatter(),
		FormatText:    TextFormatter(),
		FormatBracket: BracketFormatter(),
	}
	defaultFormatter = JSONFormatter()
)

func string2Formatter(format string) Formatter {
	if formatter, ok := formatterMap[format]; ok {
		return formatter
	}
	return defaultFormatter
}

func NewFormatterWithType(formatType FormatType) Formatter {
	return string2Formatter(formatType)
}

type Formatter interface {
	logrus.Formatter
}

func JSONFormatter() Formatter {
	return &logrus.JSONFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
	}
}

func TextFormatter() Formatter {
	return &logrus.TextFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
	}
}

// bracketFormatter
type bracketFormatter struct{}

func BracketFormatter() Formatter {
	return &bracketFormatter{}
}

func (formatter *bracketFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data[fieldKeyTime] = entry.Time.Format(timer.DefaultFormatLayout)
	data[logrus.FieldKeyMsg] = entry.Message
	data[logrus.FieldKeyLevel] = entry.Level.String()

	bracket := ""
	for k, v := range data {
		bracket = fmt.Sprintf("%s [%s:%v]", bracket, k, v)
	}
	bracket = fmt.Sprintf("%s\n", bracket)

	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}
	buffer.WriteString(strings.Trim(bracket, " "))
	return buffer.Bytes(), nil
}
