package timer

import (
	"reflect"
	"strconv"
	"time"
)

func IsStrTime(v interface{}) bool {
	t, ok := v.(string)
	if !ok {
		return false
	}
	_, err := Parse(t)
	if err != nil {
		return false
	}
	return true
}

func IsUnixTime(v interface{}) bool {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Int64, reflect.Int:
		{
			if value.Int() <= 0 {
				return false
			}
		}
	case reflect.String:
		{
			unix, err := strconv.ParseInt(value.String(), 10, 64)
			if err != nil || unix <= 0 {
				return false
			}
		}
	default:
		{
			return false
		}
	}
	return true
}

func IsTime(v interface{}) bool {
	if t, ok := v.(time.Time); ok && t.Unix() > 0 {
		return true
	}
	return false
}
