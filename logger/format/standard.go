package format

import (
	"github.com/sirupsen/logrus"
	"github.com/ihezebin/sdk/utils/timer"
)

func JSONFormatter() Formatter {
	return &logrus.JSONFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
	}
}

func TextFormatter() Formatter {
	return &logrus.TextFormatter{
		TimestampFormat: timer.DefaultFormatLayout,
	}
}
