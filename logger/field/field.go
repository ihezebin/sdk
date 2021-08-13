package field

import "github.com/sirupsen/logrus"

const (
	KeyTime      = "time"
	KeyTimestamp = "timestamp"
	KeyApp       = "app"
	KeyMsg       = "msg"
	KeyLevel     = "level"
	KeyError     = "error"
	KeyFunc      = "func"
	KeyFile      = "file"
	KeySource    = "source"
	KeyHost      = "host"
	KeyTrace     = "trace"

	ValueZero = iota
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Fields2LogrusFields convert Fields to the logrus.Fields
func Fields2LogrusFields(fields Fields) logrus.Fields {
	logrusFields := make(logrus.Fields, len(fields))
	for k, v := range fields {
		logrusFields[k] = v
	}
	return logrusFields
}
