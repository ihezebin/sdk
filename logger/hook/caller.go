package hook

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/logger/level"
	"path"
	"runtime"
	"strings"
)

/*
To create a hook, you only need to implement two methods of the hook interface：
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}
*/

const (
	kernelPackageName = "sirupsen/logrus"
	loggerPackageName = "sdk/logger"
)

// callerHook for logging the caller
type callerHook struct {
	source         bool
	simplify       bool
	skip           int
	levels         []level.Level
	fieldKeyFunc   string
	fieldKeyFile   string
	fieldKeySource string
}

// NewCallerHook Use to create the callerHook
// StandardLogger adds this Hook by default,
// which is used to increase the attributes of the specific location where the output log is located,
// so as to quickly locate the original file name and function name where the log output is located.
// StandardLogger默认添加了该Hook, 用于增加输出日志所在具体位置的属性, 便于快速定位日志输出所在的原文件名和函数名.
func NewCallerHook() *callerHook {
	return &callerHook{
		simplify:       false,
		skip:           0,
		levels:         level.AllLevels,
		fieldKeyFile:   field.KeyFile,
		fieldKeyFunc:   field.KeyFunc,
		fieldKeySource: field.KeySource,
	}
}

// Levels implement levels
func (hook *callerHook) Levels() []level.Level {
	return hook.levels
}

// Fire implement logrus.Fire
func (hook *callerHook) Fire(entry *logrus.Entry) error {
	file, function, line, err := findCaller(hook.skip)
	if err != nil {
		return err
	}
	// handle simplify
	file, function, line = hook.handleSimplify(file, function, line)
	// handle mode
	if hook.source {
		entry.Data[field.KeySource] = fmt.Sprintf("%s:%d:%s", file, line, function)
	} else {
		entry.Data[hook.fieldKeyFile] = fmt.Sprintf("%s:%d", file, line)
		entry.Data[hook.fieldKeyFunc] = function
	}
	return nil
}

// handleSimplify simplify caller info
// The full path of function name is often very long,
// so the most critical part needs to be preserved after interception
//
// 简化文件名和函数信息, 函数名的全路径往往很长, 所以需要截取后保留最关键的部分
func (hook *callerHook) handleSimplify(file string, function string, line int) (string, string, int) {
	// handle function
	if i := strings.LastIndex(function, "/"); i == -1 { // In the project root directory eg:main.main
		function = path.Ext(function)[1:]
	} else {
		function = path.Base(function)
	}
	// handle file
	if hook.simplify {
		return path.Base(file), function, line
	}
	return file, function, line
}

// Description:
// starting from the first layer of caller,
// search upward until finding the non logrus package and the sdk/logger package
// which is the interface location,
// that is the actual call location
//
// 描述：
// 从caller第一层开始，向上递进搜索, 直到找到非logrus包和非该接口所在的sdk/logger包为止, 即为实际调用位置.
//
func findCaller(skip int) (file, function string, line int, err error) {
	for i := 0; ; i++ {
		file, function, line, err = getCaller(skip + i)
		if err != nil {
			return file, function, line, err
		}
		if !strings.Contains(function, kernelPackageName) && !strings.Contains(function, loggerPackageName) {
			break
		}
	}
	return file, function, line, nil
}

// getCaller get filename, function name and line number
func getCaller(skip int) (file, function string, line int, err error) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return file, "", line, errors.New("fail to get caller, it's past the last file")
	}
	function = runtime.FuncForPC(pc).Name()
	return file, function, line, nil
}

func (hook *callerHook) SetSource(source bool) *callerHook {
	hook.source = source
	return hook
}

func (hook *callerHook) SetSimplify(simplify bool) *callerHook {
	hook.simplify = simplify
	return hook
}

func (hook *callerHook) SetLevels(levels []level.Level) *callerHook {
	hook.levels = levels
	return hook
}

func (hook *callerHook) SetFieldKeyFunc(key string) *callerHook {
	hook.fieldKeyFunc = key
	return hook
}

func (hook *callerHook) SetFieldKeyFile(key string) *callerHook {
	hook.fieldKeyFile = key
	return hook
}

func (hook *callerHook) SetFieldKeySource(key string) *callerHook {
	hook.fieldKeySource = key
	return hook
}
