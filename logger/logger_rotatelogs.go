package logger

import "time"

func RotateFileLogger(filename string, rotateTime time.Duration, expireTime time.Duration) *Logger {
	return StandardLogger().SetOutput(RotateFileOutput(filename, rotateTime, expireTime))
}
