package logger

import "github.com/sirupsen/logrus"

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// convertFields convert Fields to the logrus.Fields
func convertFields(fields Fields) logrus.Fields {
	logrusFields := make(logrus.Fields, len(fields))
	for k, v := range fields {
		logrusFields[k] = v
	}
	return logrusFields
}

type Hook = logrus.Hook

type LevelHooks = logrus.LevelHooks
