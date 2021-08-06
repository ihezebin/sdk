package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	HostKey  = "host"
	TraceKey = "trace"

	customKey = "custom"
)

var standardKeys = map[string]bool{
	logrus.FieldKeyLevel: true, // level
	logrus.FieldKeyMsg:   true, // msg
	logrus.ErrorKey:      true, // error
	FieldKeyApp:          true, // app
	FieldKeyTime:         true, // time
	FieldKeyTimestamp:    true, // timestamp
	FieldKeyFunc:         true, // func
	FieldKeyFile:         true, // file
	FieldKeySource:       true, // source
	TraceKey:             true, // trace
	HostKey:              true, // host
}

type mergeHook struct {
	levels     []Level
	ignoreKeys map[string]bool
}

// NewMergeHook Used to merge all other attributes except the standard attributes into custom,
// to avoid too many attributes and inconvenient management,
// you can add default standard attributes by passing parameters.
// 用于合并除标准属性以外的其他所有属性到custom, 避免属性过多不便管理, 可通过传参增加默认的标准属性.
func NewMergeHook(ignoreKeys ...string) *mergeHook {
	hook := &mergeHook{
		levels:     AllLevels,
		ignoreKeys: standardKeys,
	}
	for _, key := range ignoreKeys {
		hook.ignoreKeys[key] = true
	}
	return hook
}

func (hook *mergeHook) Fire(entry *logrus.Entry) error {
	filedValueCustom := map[string]interface{}{}
	for key, v := range entry.Data {
		if ok := standardKeys[key]; ok {
			continue
		}

		filedValueCustom[key] = v
		delete(entry.Data, key)
	}

	if len(filedValueCustom) > 0 {
		entry.Data[customKey] = filedValueCustom
	}

	return nil
}

func (hook *mergeHook) Levels() []Level {
	return AllLevels
}

func (hook *mergeHook) SetLevels(levels []logrus.Level) *mergeHook {
	hook.levels = levels
	return hook
}
