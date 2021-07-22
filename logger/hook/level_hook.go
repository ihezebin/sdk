package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/utils/timer"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

var (
	defaultFormatter = &logrus.JSONFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
	}
)

// WriterMap is map for mapping a log level to an io.Writer.
// Multiple levels may share a writer, but multiple writers may not be used for one level.
type WriterMap map[logrus.Level]io.Writer

// levelHook is a hook to handle writing to local log files.
type levelHook struct {
	writers   WriterMap
	levels    []logrus.Level
	lock      *sync.Mutex
	formatter logrus.Formatter

	defaultWriter io.Writer
}

// NewLevelHook returns new LFS hook.
// Output can be a string, io.Writer or WriterMap.
// If using io.Writer or WriterMap, user is responsible for closing the used io.Writer.
func NewLevelHook(output interface{}, formatter logrus.Formatter) *levelHook {
	hook := &levelHook{
		lock:   new(sync.Mutex),
		levels: logrus.AllLevels,
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
func (hook *levelHook) SetFormatter(formatter logrus.Formatter) {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	if formatter == nil {
		formatter = defaultFormatter
	}
	hook.formatter = formatter
}

// SetDefaultWriter sets default writer for levels that don't have any defined writer.
func (hook *levelHook) SetDefaultWriter(defaultWriter io.Writer) {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	hook.defaultWriter = defaultWriter
}

// Fire writes the log file to defined path or using the defined writer.
// User who run this function needs write permissions to the file or directory if the file does not yet exist.
func (hook *levelHook) Fire(entry *logrus.Entry) error {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	if hook.writers != nil || hook.defaultWriter != nil {
		return hook.write(entry)
	}
	return nil
}

// Levels returns configured log levels.
func (hook *levelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Write a log line to an io.Writer.
func (hook *levelHook) write(entry *logrus.Entry) error {
	var writer io.Writer
	if levelWriter, ok := hook.writers[entry.Level]; ok {
		writer = levelWriter
	} else {
		if hook.defaultWriter == nil {
			return nil
		}
		writer = hook.defaultWriter
	}
	// use our formatter instead of entry.String()
	msg, err := hook.formatter.Format(entry)
	if err != nil {
		log.Println("failed to format entry:", err)
		return err
	}
	_, err = writer.Write(msg)
	return err
}
