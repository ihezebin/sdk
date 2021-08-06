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
