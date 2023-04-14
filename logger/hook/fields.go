package hook

import (
	"github.com/sirupsen/logrus"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/logger/level"
)

type fieldsHook struct {
	fields field.Fields
	levels []level.Level
}

// NewFieldsHook Use to create the fieldsHook
// Used to add common log attributes
// 用于添加公共的日志属性
func NewFieldsHook(fields field.Fields) *fieldsHook {
	return &fieldsHook{
		fields: fields,
		levels: level.AllLevels,
	}
}

// Levels implement levels
func (hook *fieldsHook) Levels() []level.Level {
	return hook.levels
}

// Fire implement fire
func (hook *fieldsHook) Fire(entry *logrus.Entry) error {
	for k, v := range hook.fields {
		if k == field.KeyTimestamp {
			entry.Data[k] = entry.Time.Unix()
			continue
		}
		entry.Data[k] = v
	}
	return nil
}

func (hook *fieldsHook) SetLevels(levels []level.Level) *fieldsHook {
	hook.levels = levels
	return hook
}
