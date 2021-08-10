package command

import (
	"github.com/urfave/cli"
	"time"
)

type Command struct {
	kernel *cli.Command
}

func NewCommand(options ...CmdOption) *Command {
	return NewCommandWithConfig(newCmdConfig(options...))
}

func NewCommandWithConfig(config CmdConfig) *Command {
	return &Command{kernel: handleCommandConfig(&cli.Command{}, config)}
}

func (cmd *Command) Kernel() *cli.Command {
	if cmd.kernel == nil {
		cmd.kernel = &cli.Command{}
	}
	return cmd.kernel
}

func (cmd *Command) SubCommands() []cli.Command {
	return cmd.Kernel().Subcommands
}

func (cmd *Command) Flags() []cli.Flag {
	return cmd.Kernel().Flags
}

type Action func(v Value) error

func (cmd *Command) WithSubCommand(subCommand Command) *Command {
	cmd.Kernel().Subcommands = append(cmd.Kernel().Subcommands, *subCommand.Kernel())
	return cmd
}

func (cmd *Command) WithFlagInt(name string, value int, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.IntFlag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

func (cmd *Command) WithFlagInt64(name string, value int64, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.Int64Flag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

func (cmd *Command) WithFlagUint(name string, value uint, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.UintFlag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

func (cmd *Command) WithFlagUint64(name string, value uint64, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.Uint64Flag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

// WithFlagIntSlice example: main.ext -names Bob -names Tom -names Lisa
func (cmd *Command) WithFlagIntSlice(name string, value []int, usage string, required bool, hidden ...bool) *Command {
	intFlag := cli.IntSliceFlag{
		Name: name, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	}
	*intFlag.Value = append(*intFlag.Value, value...)
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, intFlag)
	return cmd
}

func (cmd *Command) WithFlagString(name string, value string, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.StringFlag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

func (cmd *Command) WithFlagStringSlice(name string, value []string, usage string, required bool, hidden ...bool) *Command {
	stringFlag := cli.StringSliceFlag{
		Name: name, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	}
	*stringFlag.Value = append(*stringFlag.Value, value...)
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, stringFlag)
	return cmd
}

func (cmd *Command) WithFlagBool(name string, value bool, usage string, required bool, hidden ...bool) *Command {
	if value {
		cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.BoolTFlag{
			Name: name, Usage: usage, Required: required, Hidden: len(hidden) > 0 && hidden[0],
		})
	} else {
		cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.BoolFlag{
			Name: name, Usage: usage, Required: required, Hidden: len(hidden) > 0 && hidden[0],
		})
	}
	return cmd
}

func (cmd *Command) WithFlagFloat64(name string, value float64, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.Float64Flag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}

func (cmd *Command) WithFlagDuration(name string, value time.Duration, usage string, required bool, hidden ...bool) *Command {
	cmd.Kernel().Flags = append(cmd.Kernel().Flags, cli.DurationFlag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return cmd
}
