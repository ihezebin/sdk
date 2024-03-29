package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ihezebin/sdk/logger/field"
	"github.com/ihezebin/sdk/logger/format"
	"github.com/ihezebin/sdk/logger/hook"
	"github.com/ihezebin/sdk/logger/level"
	"github.com/ihezebin/sdk/logger/writer"
	"github.com/ihezebin/sdk/utils/stringer"
	"io"
	"os"
)

type Logger struct {
	kernel *logrus.Logger
}

func New(options ...Option) *Logger {
	return NewLoggerWithConfig(newConfig(options...))
}

func NewLoggerWithConfig(config Config) *Logger {
	l := newStandardLogger()

	l.SetLevel(level.String2Level(config.Level))
	l.SetFormatter(format.String2Formatter(config.Format))

	fields := make(field.Fields)
	if stringer.NotEmpty(config.AppName) {
		fields[field.KeyApp] = config.AppName
	}
	if config.Timestamp {
		fields[field.KeyTimestamp] = field.ValueZero
	}
	l.AddHook(hook.NewFieldsHook(fields))
	if stringer.NotEmpty(config.ErrFile) {
		l.AddHook(hook.NewLocalFSHook(hook.WriterMap{
			level.ErrorLevel: FileOutput(config.ErrFile),
		}, format.String2Formatter(config.Format)))
	}

	if stringer.NotEmpty(config.File) {
		l.SetOutput(FileOutput(config.File))
	}
	if stringer.NotEmpty(config.RotateFile.File) {
		l.SetOutput(writer.NewRotateFileWriter(config.RotateFile.File, config.RotateFile.RotateTime, config.RotateFile.ExpireTime))
	}

	return l
}

func newStandardLogger() *Logger {
	return (&Logger{
		&logrus.Logger{
			Out:          StandardErrOutput(),
			Hooks:        make(logrus.LevelHooks),
			Formatter:    format.DefaultFormatter(),
			ReportCaller: false,
			Level:        level.Default(),
			ExitFunc:     os.Exit,
		},
	}).AddHook(hook.NewCallerHook().
		SetSimplify(true),
	)
}

func (logger *Logger) Kernel() *logrus.Logger {
	return logger.kernel
}

func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{logger, logger.Kernel().WithField(key, value)}
}

func (logger *Logger) WithFields(fields field.Fields) *Entry {
	return &Entry{logger, logger.Kernel().WithFields(field.Fields2LogrusFields(fields))}
}

func (logger *Logger) WithError(err error) *Entry {
	return &Entry{logger, logger.Kernel().WithError(err)}
}

func (logger *Logger) WithContext(ctx context.Context) *Entry {
	return &Entry{logger, logger.Kernel().WithContext(ctx)}
}

// Logger Printf family functions

func (logger *Logger) Logf(level level.Level, format string, args ...interface{}) {
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

func (logger *Logger) Log(level level.Level, args ...interface{}) {
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

func (logger *Logger) Logln(level level.Level, args ...interface{}) {
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
func (logger *Logger) SetLevel(level level.Level) *Logger {
	logger.Kernel().SetLevel(level)
	return logger
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() level.Level {
	return logger.Kernel().GetLevel()
}

// AddHook adds a hook to the logger hook.
func (logger *Logger) AddHook(hook hook.Hook) *Logger {
	logger.Kernel().AddHook(hook)
	return logger
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (logger *Logger) IsLevelEnabled(level level.Level) bool {
	return logger.Kernel().IsLevelEnabled(level)
}

// SetFormatter sets the logger format.
func (logger *Logger) SetFormatter(formatter format.Formatter) *Logger {
	logger.Kernel().SetFormatter(formatter)
	return logger
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(output io.Writer) *Logger {
	logger.Kernel().SetOutput(output)
	return logger
}

// DisableCaller set ReportCaller false
func (logger *Logger) DisableCaller() *Logger {
	logger.Kernel().SetReportCaller(false)
	return logger
}

// ReplaceHooks replaces the logger hook and returns the old ones
func (logger *Logger) ReplaceHooks(hooks hook.LevelHooks) hook.LevelHooks {
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
