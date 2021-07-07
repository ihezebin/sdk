package timer

import (
	"time"
)

const (
	defaultFormatLayout       = "2006-01-02 15:04:05"
	defaultFormat2MilliLayout = "2006-01-02 15:04:05.00"
)

func Now() string {
	return Format(time.Now())
}

func Format(t time.Time) string {
	return t.Format(defaultFormatLayout)
}

func Format2Milli(t time.Time) string {
	return t.Format(defaultFormat2MilliLayout)
}

func Parse(s string) (time.Time, error) {
	return time.ParseInLocation(defaultFormatLayout, s, time.Local)
}

func ParseMilli(s string) (time.Time, error) {
	return time.ParseInLocation(defaultFormat2MilliLayout, s, time.Local)
}

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}
