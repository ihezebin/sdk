package logger

func WebLogger(app string) *Logger {
	return New().SetFormatter(WebFormatter(app))
}
