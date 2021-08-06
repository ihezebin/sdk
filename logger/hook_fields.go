package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	FieldKeyTime      = "time"
	FieldKeyTimestamp = "timestamp"
	FieldKeyApp       = "app"

	fieldValueZero = iota
)

type fieldsHook struct {
	fields Fields
	levels []Level
}

// NewFieldsHook Use to create the fieldsHook
func NewFieldsHook(fields Fields) *fieldsHook {
	return &fieldsHook{
		fields: fields,
		levels: AllLevels,
	}
}

// Levels implement levels
func (hook *fieldsHook) Levels() []Level {
	return hook.levels
}

// Fire implement fire
func (hook *fieldsHook) Fire(entry *logrus.Entry) error {
	for k, v := range hook.fields {
		if k == FieldKeyTimestamp {
			entry.Data[k] = entry.Time.Unix()
			continue
		}
		entry.Data[k] = v
	}
	return nil
}

func (hook *fieldsHook) SetLevels(levels []logrus.Level) *fieldsHook {
	hook.levels = levels
	return hook
}
