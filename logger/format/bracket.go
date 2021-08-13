package format

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk/logger/field"
	"github.com/whereabouts/sdk/utils/timer"
	"strings"
)

// bracketFormatter
type bracketFormatter struct{}

func BracketFormatter() Formatter {
	return &bracketFormatter{}
}

// Format implement logrus.Formatter
func (formatter *bracketFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(field.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data[field.KeyTime] = entry.Time.Format(timer.DefaultFormatLayout)
	data[field.KeyMsg] = entry.Message
	data[field.KeyLevel] = entry.Level.String()

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
