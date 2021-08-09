package cli

import (
	"github.com/urfave/cli"
	"time"
)

type Flag interface {
	Int(name string) int
	GlobalInt(name string) int
	Int64(name string) int64
	GlobalInt64(name string) int64
	Uint(name string) uint
	GlobalUint(name string) uint
	Uint64(name string) uint64
	GlobalUint64(name string) uint64
	IntSlice(name string) []int
	GlobalIntSlice(name string) []int
	Int64Slice(name string) []int64
	GlobalInt64Slice(name string) []int64
	String(name string) string
	GlobalString(name string) string
	StringSlice(name string) []string
	GlobalStringSlice(name string) []string
	Duration(name string) time.Duration
	GlobalDuration(name string) time.Duration
	Bool(name string) bool
	GlobalBool(name string) bool
	NumFlags() int
	Set(name, value string) error
	GlobalSet(name, value string) error
	IsSet(name string) bool
	GlobalIsSet(name string) bool
	FlagNames() (names []string)
	GlobalFlagNames() (names []string)
}

type flag struct {
	app *App
	ctx *cli.Context
}

// Int looks up the value of a local IntFlag, returns
// 0 if not found
func (f *flag) Int(name string) int {
	return f.ctx.Int(name)
}

// GlobalInt looks up the value of a global IntFlag, returns
// 0 if not found
func (f *flag) GlobalInt(name string) int {
	return f.ctx.GlobalInt(name)
}

// Int64 looks up the value of a local Int64Flag, returns
// 0 if not found
func (f *flag) Int64(name string) int64 {
	return f.ctx.Int64(name)
}

// GlobalInt64 looks up the value of a global Int64Flag, returns
// 0 if not found
func (f *flag) GlobalInt64(name string) int64 {
	return f.ctx.GlobalInt64(name)
}

// Uint looks up the value of a local UintFlag, returns
// 0 if not found
func (f *flag) Uint(name string) uint {
	return f.ctx.Uint(name)
}

// GlobalUint looks up the value of a global UintFlag, returns
// 0 if not found
func (f *flag) GlobalUint(name string) uint {
	return f.ctx.GlobalUint(name)
}

// Uint64 looks up the value of a local Uint64Flag, returns
// 0 if not found
func (f *flag) Uint64(name string) uint64 {
	return f.ctx.Uint64(name)
}

// GlobalUint64 looks up the value of a global Uint64Flag, returns
// 0 if not found
func (f *flag) GlobalUint64(name string) uint64 {
	return f.ctx.GlobalUint64(name)
}

// IntSlice looks up the value of a local IntSliceFlag, returns
// nil if not found
func (f *flag) IntSlice(name string) []int {
	return f.ctx.IntSlice(name)
}

// GlobalIntSlice looks up the value of a global IntSliceFlag, returns
// nil if not found
func (f *flag) GlobalIntSlice(name string) []int {
	return f.ctx.GlobalIntSlice(name)
}

// Int64Slice looks up the value of a local Int64SliceFlag, returns
// nil if not found
func (f *flag) Int64Slice(name string) []int64 {
	return f.ctx.Int64Slice(name)
}

// GlobalInt64Slice looks up the value of a global Int64SliceFlag, returns
// nil if not found
func (f *flag) GlobalInt64Slice(name string) []int64 {
	return f.ctx.GlobalInt64Slice(name)
}

// String looks up the value of a local StringFlag, returns
// "" if not found
func (f *flag) String(name string) string {
	return f.ctx.String(name)
}

// GlobalString looks up the value of a global StringFlag, returns
// "" if not found
func (f *flag) GlobalString(name string) string {
	return f.ctx.GlobalString(name)
}

// StringSlice looks up the value of a local StringSliceFlag, returns
// nil if not found
func (f *flag) StringSlice(name string) []string {
	return f.ctx.StringSlice(name)
}

// GlobalStringSlice looks up the value of a global StringSliceFlag, returns
// nil if not found
func (f *flag) GlobalStringSlice(name string) []string {
	return f.ctx.GlobalStringSlice(name)
}

// Duration looks up the value of a local DurationFlag, returns
// 0 if not found
func (f *flag) Duration(name string) time.Duration {
	return f.ctx.Duration(name)
}

// GlobalDuration looks up the value of a global DurationFlag, returns
// 0 if not found
func (f *flag) GlobalDuration(name string) time.Duration {
	return f.ctx.GlobalDuration(name)
}

func (f *flag) Bool(name string) bool {
	return f.ctx.Bool(name)
}

// GlobalBool looks up the value of a global BoolFlag, returns
// false if not found
func (f *flag) GlobalBool(name string) bool {
	return f.ctx.GlobalBool(name)
}

// NumFlags returns the number of flags set
func (f *flag) NumFlags() int {
	return f.ctx.NumFlags()
}

// Set sets a context flag to a value.
func (f *flag) Set(name, value string) error {
	return f.ctx.Set(name, value)
}

// GlobalSet sets a context flag to a value on the global flagset
func (f *flag) GlobalSet(name, value string) error {
	return f.ctx.GlobalSet(name, value)
}

// IsSet determines if the flag was actually set
func (f *flag) IsSet(name string) bool {
	return f.ctx.IsSet(name)
}

// GlobalIsSet determines if the global flag was actually set
func (f *flag) GlobalIsSet(name string) bool {
	return f.ctx.GlobalIsSet(name)
}

// FlagNames returns a slice of flag names used in this context.
func (f *flag) FlagNames() (names []string) {
	return f.ctx.FlagNames()
}

// GlobalFlagNames returns a slice of global flag names used by the app.
func (f *flag) GlobalFlagNames() (names []string) {
	return f.ctx.GlobalFlagNames()
}

// Args contains apps console arguments
type Args = cli.Args

// NArg returns the number of the command line arguments.
func (f *flag) NArg() int {
	return f.ctx.NArg()
}

// Args returns the command line arguments associated with the context.
func (f *flag) Args() Args {
	return f.ctx.Args()
}
