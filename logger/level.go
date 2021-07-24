package logger

import (
	"github.com/sirupsen/logrus"
)

// Level type
type Level = logrus.Level

var defaultLevel = InfoLevel

// AllLevels A constant exposing all logging levels
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hook to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel

	PanicLevelString   = "panic"
	FatalLevelString   = "fatal"
	ErrorLevelString   = "error"
	WarnLevelString    = "warn"
	WarningLevelString = "warning"
	InfoLevelString    = "info"
	DebugLevelString   = "debug"
	TraceLevelString   = "trace"
)

var levelMap = map[string]Level{
	"panic":   PanicLevel,
	"fatal":   FatalLevel,
	"error":   ErrorLevel,
	"warn":    WarnLevel,
	"warning": WarnLevel,
	"info":    InfoLevel,
	"debug":   DebugLevel,
	"trace":   TraceLevel,
}

func string2Level(levelStr string) Level {
	if level, ok := levelMap[levelStr]; ok {
		return level
	}
	return defaultLevel
}
