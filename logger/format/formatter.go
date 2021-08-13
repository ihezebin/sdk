package format

import (
	"github.com/sirupsen/logrus"
)

type Type = string

const (
	JSON    Type = "json"
	Text    Type = "text"
	Bracket Type = "bracket"
)

var (
	formatterMap = map[Type]Formatter{
		JSON:    JSONFormatter(),
		Text:    TextFormatter(),
		Bracket: BracketFormatter(),
	}
	defaultFormatter = JSONFormatter()
)

type Formatter interface {
	logrus.Formatter
}

func DefaultFormatter() Formatter {
	return defaultFormatter
}

func String2Formatter(format string) Formatter {
	if formatter, ok := formatterMap[format]; ok {
		return formatter
	}
	return defaultFormatter
}

func NewFormatterWithType(formatType Type) Formatter {
	return String2Formatter(formatType)
}
