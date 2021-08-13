package hook

import (
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk/logger/field"
	"github.com/whereabouts/sdk/logger/level"
)

const (
	customKey = "custom"
)

var standardKeys = map[string]bool{
	field.KeyLevel:     true, // level
	field.KeyMsg:       true, // msg
	field.KeyError:     true, // error
	field.KeyApp:       true, // app
	field.KeyTime:      true, // time
	field.KeyTimestamp: true, // timestamp
	field.KeyFunc:      true, // func
	field.KeyFile:      true, // file
	field.KeySource:    true, // source
	field.KeyTrace:     true, // trace
	field.KeyHost:      true, // host
}

type mergeHook struct {
	levels     []level.Level
	ignoreKeys map[string]bool
}

// NewMergeHook Used to merge all other attributes except the standard attributes into custom,
// to avoid too many attributes and inconvenient management,
// you can add default standard attributes by passing parameters.
// 用于合并除标准属性以外的其他所有属性到custom, 避免属性过多不便管理, 可通过传参增加默认的标准属性.
func NewMergeHook(ignoreKeys ...string) *mergeHook {
	hook := &mergeHook{
		levels:     level.AllLevels,
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

func (hook *mergeHook) Levels() []level.Level {
	return level.AllLevels
}

func (hook *mergeHook) SetLevels(levels []logrus.Level) *mergeHook {
	hook.levels = levels
	return hook
}
