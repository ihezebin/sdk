package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	fieldKeyTime      = "time"
	fieldKeyTimestamp = "timestamp"
	fieldKeyApp       = "app"
	fieldValueZero    = iota
)

type fieldsHook struct {
	fields Fields
	levels []Level
}

// NewFieldsHook Use to create the fieldsHook
func NewFieldsHook(fields Fields) *fieldsHook {
	hook := fieldsHook{
		fields: fields,
		levels: AllLevels,
	}
	return &hook
}

// Levels implement levels
func (hook *fieldsHook) Levels() []Level {
	return hook.levels
}

// Fire implement fire
func (hook *fieldsHook) Fire(entry *logrus.Entry) error {
	for k, v := range hook.fields {
		if k == fieldKeyTimestamp {
			entry.Data[k] = entry.Time.Unix()
			continue
		}
		entry.Data[k] = v
	}
	return nil
}
