package command

import (
	"github.com/urfave/cli"
	"time"
)

type Value interface {
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

type value struct {
	ctx *cli.Context
}

func NewValue(ctx *cli.Context) Value {
	return &value{ctx: ctx}
}

// Int looks up the value of a local IntFlag, returns
// 0 if not found
func (v *value) Int(name string) int {
	return v.ctx.Int(name)
}

// GlobalInt looks up the value of a global IntFlag, returns
// 0 if not found
func (v *value) GlobalInt(name string) int {
	return v.ctx.GlobalInt(name)
}

// Int64 looks up the value of a local Int64Flag, returns
// 0 if not found
func (v *value) Int64(name string) int64 {
	return v.ctx.Int64(name)
}

// GlobalInt64 looks up the value of a global Int64Flag, returns
// 0 if not found
func (v *value) GlobalInt64(name string) int64 {
	return v.ctx.GlobalInt64(name)
}

// Uint looks up the value of a local UintFlag, returns
// 0 if not found
func (v *value) Uint(name string) uint {
	return v.ctx.Uint(name)
}

// GlobalUint looks up the value of a global UintFlag, returns
// 0 if not found
func (v *value) GlobalUint(name string) uint {
	return v.ctx.GlobalUint(name)
}

// Uint64 looks up the value of a local Uint64Flag, returns
// 0 if not found
func (v *value) Uint64(name string) uint64 {
	return v.ctx.Uint64(name)
}

// GlobalUint64 looks up the value of a global Uint64Flag, returns
// 0 if not found
func (v *value) GlobalUint64(name string) uint64 {
	return v.ctx.GlobalUint64(name)
}

// IntSlice looks up the value of a local IntSliceFlag, returns
// nil if not found
func (v *value) IntSlice(name string) []int {
	return v.ctx.IntSlice(name)
}

// GlobalIntSlice looks up the value of a global IntSliceFlag, returns
// nil if not found
func (v *value) GlobalIntSlice(name string) []int {
	return v.ctx.GlobalIntSlice(name)
}

// Int64Slice looks up the value of a local Int64SliceFlag, returns
// nil if not found
func (v *value) Int64Slice(name string) []int64 {
	return v.ctx.Int64Slice(name)
}

// GlobalInt64Slice looks up the value of a global Int64SliceFlag, returns
// nil if not found
func (v *value) GlobalInt64Slice(name string) []int64 {
	return v.ctx.GlobalInt64Slice(name)
}

// String looks up the value of a local StringFlag, returns
// "" if not found
func (v *value) String(name string) string {
	return v.ctx.String(name)
}

// GlobalString looks up the value of a global StringFlag, returns
// "" if not found
func (v *value) GlobalString(name string) string {
	return v.ctx.GlobalString(name)
}

// StringSlice looks up the value of a local StringSliceFlag, returns
// nil if not found
func (v *value) StringSlice(name string) []string {
	return v.ctx.StringSlice(name)
}

// GlobalStringSlice looks up the value of a global StringSliceFlag, returns
// nil if not found
func (v *value) GlobalStringSlice(name string) []string {
	return v.ctx.GlobalStringSlice(name)
}

// Duration looks up the value of a local DurationFlag, returns
// 0 if not found
func (v *value) Duration(name string) time.Duration {
	return v.ctx.Duration(name)
}

// GlobalDuration looks up the value of a global DurationFlag, returns
// 0 if not found
func (v *value) GlobalDuration(name string) time.Duration {
	return v.ctx.GlobalDuration(name)
}

func (v *value) Bool(name string) bool {
	return v.ctx.Bool(name)
}

// GlobalBool looks up the value of a global BoolFlag, returns
// false if not found
func (v *value) GlobalBool(name string) bool {
	return v.ctx.GlobalBool(name)
}

// NumFlags returns the number of flags set
func (v *value) NumFlags() int {
	return v.ctx.NumFlags()
}

// Set sets a context flag to a value.
func (v *value) Set(name, value string) error {
	return v.ctx.Set(name, value)
}

// GlobalSet sets a context flag to a value on the global flagset
func (v *value) GlobalSet(name, value string) error {
	return v.ctx.GlobalSet(name, value)
}

// IsSet determines if the flag was actually set
func (v *value) IsSet(name string) bool {
	return v.ctx.IsSet(name)
}

// GlobalIsSet determines if the global flag was actually set
func (v *value) GlobalIsSet(name string) bool {
	return v.ctx.GlobalIsSet(name)
}

// FlagNames returns a slice of flag names used in this context.
func (v *value) FlagNames() (names []string) {
	return v.ctx.FlagNames()
}

// GlobalFlagNames returns a slice of global flag names used by the app.
func (v *value) GlobalFlagNames() (names []string) {
	return v.ctx.GlobalFlagNames()
}

// Args contains apps console arguments
type Args = cli.Args

// NArg returns the number of the command line arguments.
func (v *value) NArg() int {
	return v.ctx.NArg()
}

// Args returns the command line arguments associated with the context.
func (v *value) Args() Args {
	return v.ctx.Args()
}
