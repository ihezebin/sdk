package cli

import (
	"time"
)

var gApp = NewApp(WithHideHelp(true)).SetDefaultHelp(false)

func Run() (value Value, err error) {
	err = gApp.SetAction(func(v Value) error {
		value = v
		return nil
	}).Run()
	return
}

func WithFlagInt(name string, value int, usage string) {
	gApp.WithFlagInt(name, value, usage, false)
}

func WithFlagInt64(name string, value int64, usage string) {
	gApp.WithFlagInt64(name, value, usage, false)
}

func WithFlagUint(name string, value uint, usage string) {
	gApp.WithFlagUint(name, value, usage, false)
}

func WithFlagUint64(name string, value uint64, usage string) {
	gApp.WithFlagUint64(name, value, usage, false)
}

func WithFlagIntSlice(name string, value []int, usage string) {
	gApp.WithFlagIntSlice(name, value, usage, false)
}

func WithFlagString(name string, value string, usage string) {
	gApp.WithFlagString(name, value, usage, false)
}

// WithFlagStringSlice example: main.ext -names Bob -names Tom -names Lisa
func WithFlagStringSlice(name string, value []string, usage string) {
	gApp.WithFlagStringSlice(name, value, usage, false)
}

func WithFlagBool(name string, value bool, usage string) {
	gApp.WithFlagBool(name, value, usage, false)
}

func WithFlagFloat64(name string, value float64, usage string) {
	gApp.WithFlagFloat64(name, value, usage, false)
}

func WithFlagDuration(name string, value time.Duration, usage string) {
	gApp.WithFlagDuration(name, value, usage, false)
}
