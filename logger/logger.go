package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger struct {
	kernel *logrus.Logger
}

func New() *Logger {
	return (&Logger{
		&logrus.Logger{
			Out:          StandardErrOutput(),
			Hooks:        make(logrus.LevelHooks),
			Formatter:    JSONFormatter(),
			ReportCaller: false,
			Level:        InfoLevel,
			ExitFunc:     os.Exit,
		},
	}).AddHook(NewCallerHook().SetSimplify(true))
}

func (logger *Logger) Kernel() *logrus.Logger {
	return logger.kernel
}

func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{logger, logger.Kernel().WithField(key, value)}
}

func (logger *Logger) WithFields(fields Fields) *Entry {
	return &Entry{logger, logger.Kernel().WithFields(fields2LogrusFields(fields))}
}

func (logger *Logger) WithError(err error) *Entry {
	return &Entry{logger, logger.Kernel().WithError(err)}
}

func (logger *Logger) WithContext(ctx context.Context) *Entry {
	return &Entry{logger, logger.Kernel().WithContext(ctx)}
}

// Logger Printf family functions

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	logger.Kernel().Logf(level, format, args...)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.Kernel().Tracef(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Kernel().Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Kernel().Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.Kernel().Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Kernel().Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.Kernel().Warningf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Kernel().Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Kernel().Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.Kernel().Panicf(format, args...)
}

// Logger Print family functions

func (logger *Logger) Log(level Level, args ...interface{}) {
	logger.Kernel().Log(level, args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.Kernel().Trace(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Kernel().Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Kernel().Info(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.Kernel().Print(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Kernel().Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.Kernel().Warning(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Kernel().Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.Kernel().Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.Kernel().Panic(args...)
}

// Logger Println family functions

func (logger *Logger) Logln(level Level, args ...interface{}) {
	logger.Kernel().Logln(level, args...)
}

func (logger *Logger) Traceln(args ...interface{}) {
	logger.Kernel().Traceln(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.Kernel().Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.Kernel().Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.Kernel().Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.Kernel().Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	logger.Kernel().Warningln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.Kernel().Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.Kernel().Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.Kernel().Panicln(args...)
}

// SetLevel sets the logger level.
func (logger *Logger) SetLevel(level Level) *Logger {
	logger.Kernel().SetLevel(level)
	return logger
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() Level {
	return logger.Kernel().GetLevel()
}

// AddHook adds a hook to the logger hook.
func (logger *Logger) AddHook(hook Hook) *Logger {
	logger.Kernel().AddHook(hook)
	return logger
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (logger *Logger) IsLevelEnabled(level Level) bool {
	return logger.Kernel().IsLevelEnabled(level)
}

// SetFormatter sets the logger formatter.
func (logger *Logger) SetFormatter(formatter Formatter) *Logger {
	logger.Kernel().SetFormatter(formatter)
	return logger
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(output io.Writer) *Logger {
	logger.Kernel().SetOutput(output)
	return logger
}

// CloseCaller set ReportCaller false
func (logger *Logger) DisableCaller() *Logger {
	logger.Kernel().SetReportCaller(false)
	return logger
}

// ReplaceHooks replaces the logger hook and returns the old ones
func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks {
	return logger.Kernel().ReplaceHooks(hooks)
}

// SetNoLock
// When file is opened with appending mode, it's safe to
// write concurrently to a file (within 4k message on Linux).
// In these cases user can choose to disable the lock.
func (logger *Logger) SetNoLock() *Logger {
	logger.Kernel().SetNoLock()
	return logger
}
