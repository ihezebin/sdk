package logger

func RotateFileLogger() *Logger {
	return New().SetOutput(RotateFileOutput())
}
