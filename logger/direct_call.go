package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/whereabouts/sdk-go/logger/hooks"
	"io"
	"time"
)

var (
	defaultLogger = New()
	lineHook      = hooks.NewLineHook("file")
)

// The default output format is JSON, And the default format time is "2006-01-02 15:04:05"
// Add file line number and function by default
func init() {
	defaultLogger.l.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	defaultLogger.l.AddHook(lineHook)
}

func DefaultLogger() *Logger {
	return defaultLogger
}

// IsShowLine setting if show the file line number and function
func SetLine(line bool) {
	if !line {
		levelHooks := defaultLogger.l.Hooks
		for level, hooks := range levelHooks {
			levelHooks[level] = deleteHook(hooks, lineHook)
		}
		defaultLogger.l.ReplaceHooks(levelHooks)
	}
}

func deleteHook(s []logrus.Hook, elem interface{}) []logrus.Hook {
	for i, v := range s {
		if v == elem {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// SetFormatter sets the default logger output
func SetOutput(out io.Writer) {
	defaultLogger.l.SetOutput(out)
}

// SetFormatter sets the default logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	defaultLogger.l.SetFormatter(formatter)
}

// SetLevel sets the default logger level.
func SetLevel(level logrus.Level) {
	defaultLogger.l.SetLevel(level)
}

// GetLevel returns the default logger level.
func GetLevel() logrus.Level {
	return defaultLogger.l.GetLevel()
}

// IsLevelEnabled checks if the log level of the default logger is greater than the level param
func IsLevelEnabled(level logrus.Level) bool {
	return defaultLogger.l.IsLevelEnabled(level)
}

// AddHook adds a hook to the default logger hooks.
func AddHook(hook logrus.Hook) {
	defaultLogger.l.AddHook(hook)
}

// WithError creates an entry from the default logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return defaultLogger.l.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the default logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return defaultLogger.l.WithContext(ctx)
}

// WithField creates an entry from the default logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return defaultLogger.l.WithField(key, value)
}

// WithFields creates an entry from the default logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return defaultLogger.l.WithFields(fields)
}

// WithTime creats an entry from the default logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return defaultLogger.l.WithTime(t)
}

// Trace logs a message at level Trace on the default logger.
func Trace(args ...interface{}) {
	defaultLogger.l.Trace(args...)
}

// Debug logs a message at level Debug on the default logger.
func Debug(args ...interface{}) {
	defaultLogger.l.Debug(args...)
}

// Print logs a message at level Info on the default logger.
func Print(args ...interface{}) {
	defaultLogger.l.Print(args...)
}

// Info logs a message at level Info on the default logger.
func Info(args ...interface{}) {
	defaultLogger.l.Info(args...)
}

// Warn logs a message at level Warn on the default logger.
func Warn(args ...interface{}) {
	defaultLogger.l.Warn(args...)
}

// Warning logs a message at level Warn on the default logger.
func Warning(args ...interface{}) {
	defaultLogger.l.Warning(args...)
}

// Error logs a message at level Error on the default logger.
func Error(args ...interface{}) {
	defaultLogger.l.Error(args...)
}

// Panic logs a message at level Panic on the default logger.
func Panic(args ...interface{}) {
	defaultLogger.l.Panic(args...)
}

// Fatal logs a message at level Fatal on the default logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	defaultLogger.l.Fatal(args...)
}

// Tracef logs a message at level Trace on the default logger.
func Tracef(format string, args ...interface{}) {
	defaultLogger.l.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the default logger.
func Debugf(format string, args ...interface{}) {
	defaultLogger.l.Debugf(format, args...)
}

// Printf logs a message at level Info on the default logger.
func Printf(format string, args ...interface{}) {
	defaultLogger.l.Printf(format, args...)
}

// Infof logs a message at level Info on the default logger.
func Infof(format string, args ...interface{}) {
	defaultLogger.l.Infof(format, args...)
}

// Warnf logs a message at level Warn on the default logger.
func Warnf(format string, args ...interface{}) {
	defaultLogger.l.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the default logger.
func Warningf(format string, args ...interface{}) {
	defaultLogger.l.Warningf(format, args...)
}

// Errorf logs a message at level Error on the default logger.
func Errorf(format string, args ...interface{}) {
	defaultLogger.l.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the default logger.
func Panicf(format string, args ...interface{}) {
	defaultLogger.l.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the default logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.l.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the default logger.
func Traceln(args ...interface{}) {
	defaultLogger.l.Traceln(args...)
}

// Debugln logs a message at level Debug on the default logger.
func Debugln(args ...interface{}) {
	defaultLogger.l.Debugln(args...)
}

// Println logs a message at level Info on the default logger.
func Println(args ...interface{}) {
	defaultLogger.l.Println(args...)
}

// Infoln logs a message at level Info on the default logger.
func Infoln(args ...interface{}) {
	defaultLogger.l.Infoln(args...)
}

// Warnln logs a message at level Warn on the default logger.
func Warnln(args ...interface{}) {
	defaultLogger.l.Warnln(args...)
}

// Warningln logs a message at level Warn on the default logger.
func Warningln(args ...interface{}) {
	defaultLogger.l.Warningln(args...)
}

// Errorln logs a message at level Error on the default logger.
func Errorln(args ...interface{}) {
	defaultLogger.l.Errorln(args...)
}

// Panicln logs a message at level Panic on the default logger.
func Panicln(args ...interface{}) {
	defaultLogger.l.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the default logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	defaultLogger.l.Fatalln(args...)
}
