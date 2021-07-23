package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Entry struct {
	logger *Logger
	kernel *logrus.Entry
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		logger: logger,
		kernel: logrus.NewEntry(logger.Kernel()),
	}
}

func (entry *Entry) logrusEntry2Entry(logrusEntry *logrus.Entry) *Entry {
	return &Entry{entry.logger, logrusEntry}
}

func (entry *Entry) Kernel() *logrus.Entry {
	return entry.kernel
}

// Bytes Returns the bytes representation of this entry from the formatter.
func (entry *Entry) Bytes() ([]byte, error) {
	return entry.Kernel().Bytes()
}

// String Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	return entry.Kernel().String()
}

// WithError Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return entry.logrusEntry2Entry(entry.Kernel().WithError(err))
}

// WithContext Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	return entry.logrusEntry2Entry(entry.Kernel().WithContext(ctx))
}

// WithField Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.logrusEntry2Entry(entry.Kernel().WithField(key, value))
}

// WithFields Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	return entry.logrusEntry2Entry(entry.Kernel().WithFields(fields2LogrusFields(fields)))
}

func (entry Entry) HasCaller() bool {
	return entry.Kernel().HasCaller()
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	entry.Kernel().Log(level, args...)
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.Kernel().Trace(args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.Kernel().Debug(args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Kernel().Print(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.Kernel().Info(args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.Kernel().Warn(args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.Kernel().Warning(args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.Kernel().Error(args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.Kernel().Fatal(args...)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.Kernel().Panic(args...)
}

// Entry Printf family functions

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	entry.Kernel().Logf(level, format, args...)
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.Kernel().Tracef(format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.Kernel().Debugf(format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.Kernel().Infof(format, args...)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.Kernel().Printf(format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.Kernel().Warnf(format, args...)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.Kernel().Warningf(format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.Kernel().Errorf(format, args...)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.Kernel().Fatalf(format, args...)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.Kernel().Panicf(format, args...)
}

// Entry Println family functions

func (entry *Entry) Logln(level Level, args ...interface{}) {
	entry.Kernel().Logln(level, args...)
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.Kernel().Traceln(args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.Kernel().Debugln(args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.Kernel().Infoln(args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.Kernel().Println(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.Kernel().Warnln(args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.Kernel().Warningln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.Kernel().Errorln(args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.Kernel().Fatalln(args...)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.Kernel().Panicln(args...)
}
