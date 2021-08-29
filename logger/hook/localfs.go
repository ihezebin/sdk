package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/sdk/logger/format"
	"github.com/whereabouts/sdk/logger/level"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

// WriterMap is map for mapping a log level to an io.Writer.
// Multiple levels may share a writer, but multiple writers may not be used for one level.
type WriterMap map[level.Level]io.Writer

// levelHook is a hook to handle writing to local log files.
type localFSHook struct {
	writers   WriterMap
	levels    []level.Level
	lock      *sync.Mutex
	formatter format.Formatter

	defaultWriter io.Writer
}

// NewLocalFSHook returns new LFS hook.
// Output can be a string, io.Writer or WriterMap.
// If using io.Writer or WriterMap, user is responsible for closing the used io.Writer.
// 日志内容分级别输出到本地文件系统, 如: INFO级别输出到test.log文件, ERROR级别输出到test.err.log文件, DEBUG级别输出到os.Stderr
func NewLocalFSHook(output interface{}, formatter format.Formatter) *localFSHook {
	hook := &localFSHook{
		lock:   new(sync.Mutex),
		levels: level.AllLevels,
	}
	hook.SetFormatter(formatter)
	switch output.(type) {
	case string:
		filename := output.(string)
		dir := filepath.Dir(filename)
		_ = os.MkdirAll(dir, os.ModePerm)
		writer, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(fmt.Sprintf("create file writer failed: %s", err.Error()))
		}
		hook.SetDefaultWriter(writer)
		break
	case io.Writer:
		hook.SetDefaultWriter(output.(io.Writer))
		break
	case WriterMap:
		hook.writers = output.(WriterMap)
		break
	default:
		panic(fmt.Sprintf("unsupported output type: %v", reflect.TypeOf(output)))
	}
	return hook
}

// SetFormatter sets the format that will be used by hook.
func (hook *localFSHook) SetFormatter(formatter format.Formatter) {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	if formatter == nil {
		formatter = format.DefaultFormatter()
	}
	hook.formatter = formatter
}

// SetDefaultWriter sets default writer for levels that don't have any defined writer.
func (hook *localFSHook) SetDefaultWriter(defaultWriter io.Writer) {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	hook.defaultWriter = defaultWriter
}

// Fire writes the log file to defined path or using the defined writer.
// User who run this function needs write permissions to the file or directory if the file does not yet exist.
func (hook *localFSHook) Fire(entry *logrus.Entry) error {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	if hook.writers != nil || hook.defaultWriter != nil {
		return hook.write(entry)
	}
	return nil
}

// Levels returns configured log levels.
func (hook *localFSHook) Levels() []level.Level {
	return hook.levels
}

// Write a log line to an io.Writer.
func (hook *localFSHook) write(entry *logrus.Entry) error {
	var writer io.Writer
	if levelWriter, ok := hook.writers[entry.Level]; ok {
		writer = levelWriter
	} else {
		if hook.defaultWriter == nil {
			return nil
		}
		writer = hook.defaultWriter
	}
	// use our format instead of entry.String()
	msg, err := hook.formatter.Format(entry)
	if err != nil {
		logger.Println("failed to format entry:", err)
		return err
	}
	_, err = writer.Write(msg)
	return err
}
