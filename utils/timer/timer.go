package timer

import (
	"time"
)

const (
	DefaultFormatLayout       = "2006-01-02 15:04:05"
	DefaultFormat2MilliLayout = "2006-01-02 15:04:05.00"
)

func Now() string {
	return Format(time.Now())
}

func Format(t time.Time) string {
	return t.Format(DefaultFormatLayout)
}

func Format2Milli(t time.Time) string {
	return t.Format(DefaultFormat2MilliLayout)
}

func Parse(s string) (time.Time, error) {
	return time.ParseInLocation(DefaultFormatLayout, s, time.Local)
}

func ParseMilli(s string) (time.Time, error) {
	return time.ParseInLocation(DefaultFormat2MilliLayout, s, time.Local)
}

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}
