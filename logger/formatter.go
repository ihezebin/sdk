package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/utils/timer"
)

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

func BracketFormatter() Formatter {
	return &bracketFormatter{}
}

func WebFormatter(app string) Formatter {
	return &webFormatter{app}
}

// bracketFormatter
type bracketFormatter struct{}

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
	data[logrus.FieldKeyTime] = entry.Time.Format(timer.DefaultFormatLayout)
	data[logrus.FieldKeyMsg] = entry.Message
	data[logrus.FieldKeyLevel] = entry.Level.String()

	log := ""
	for k, v := range data {
		log = fmt.Sprintf("%s [%s:%v]", log, k, v)
	}
	log = fmt.Sprintf("%s\n", log)

	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}
	buffer.WriteString(log[1:])
	return buffer.Bytes(), nil
}

// webFormatter
type webFormatter struct {
	app string
}

const (
	fieldKeyTimestamp = "timestamp"
	fieldKeyApp       = "app"
)

func (formatter *webFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+5)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data[logrus.FieldKeyTime] = entry.Time.Format(timer.DefaultFormatLayout)
	data[fieldKeyTimestamp] = entry.Time.Unix()
	data[fieldKeyApp] = formatter.app
	data[logrus.FieldKeyMsg] = entry.Message
	data[logrus.FieldKeyLevel] = entry.Level.String()

	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}
	return buffer.Bytes(), nil
}
