package hooks

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

/*
	To create a hooks, you only need to implement two methods of the hooks interface：
	type Hook interface {
		Levels() []Level
		Fire(*Entry) error
	}
*/

// lineHook for logging the caller
type LineHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// Use to create an hooks
// Please provide the field name of the line, and optional log level
func NewLineHook(field string, levels ...logrus.Level) logrus.Hook {
	hook := LineHook{
		Field:  field,
		Skip:   0,
		levels: levels,
	}
	if field == "" {
		hook.Field = "line"
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook LineHook) Levels() []logrus.Level {
	return hook.levels
}

// Fire implement fire
func (hook LineHook) Fire(entry *logrus.Entry) error {
	var err error
	// add the line info to field
	entry.Data[hook.Field], entry.Data["func"], err = findCaller(hook.Skip)
	return err
}

// Description:
// starting from the first layer of caller,
// search upward until finding the non logrus package and the logger package
// which is the interface location,
// that is the actual call location
//
// 描述：
// 从caller第一层开始，向上递进搜索, 直到找到非logrus包和非该接口所在的logger包为止，即为实际调用位置.
//
func findCaller(skip int) (file, fn string, err error) {
	var line int
	for i := 0; ; i++ {
		file, fn, line, err = getCaller(skip + i)
		if err != nil {
			return file, fn, err
		}
		if !strings.Contains(file, "logrus") && !strings.Contains(file, "logger") {
			file = fmt.Sprintf("%s:%d", file, line)
			break
		}
	}
	return file, fn, nil
}

// Description:
// The full path of a file is often very long,
// so it is necessary to keep the most critical part after being intercepted,
// But because the file name is often repeated in multiple packages,
// we choose to take one more layer and get to the package name where the file is located
//
// 描述：
// 文件的全路径往往很长, 所以需要截取后保留最关键的部分,
// 但又由于文件名在多个包中往往有重复, 因此选择多取一层, 取到文件所在的包名处.
//
func getCaller(skip int) (file, fn string, line int, err error) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return file, "", line, errors.New("fail to get caller ,it's past the last file")
	}
	fn = runtime.FuncForPC(pc).Name()
	//fmt.Println(file, fn, line)
	for i := strings.LastIndex(fn, "/"); i < len(fn); i++ {
		// In the project root directory eg:main.main
		if i == -1 {
			file = file[strings.LastIndex(file, "/")+1:]
			fn = fn[strings.Index(fn, ".")+1:]
			break
		} else if fn[i] == '.' {
			file = fmt.Sprintf("%s/%s", fn[:i], file[strings.LastIndex(file, "/")+1:])
			fn = fn[i+1:]
			break
		}
	}
	return file, fn, line, nil
}
