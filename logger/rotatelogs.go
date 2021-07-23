package logger

func RotateFileLogger(filename string) *Logger {
	return New().SetOutput(RotateFileOutput(filename))
}
