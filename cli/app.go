package cli

import (
	"github.com/urfave/cli"
	"github.com/ihezebin/sdk/cli/command"
	"github.com/ihezebin/sdk/cli/flag"
	"os"
	"sort"
	"time"
)

var (
	NilAction = func(v Value) error {
		return nil
	}
	HelpAction = func(v Value) error {
		args := v.Args()
		if args.Present() {
			return cli.ShowCommandHelp(v.Kernel(), args.First())
		}
		_ = cli.ShowAppHelp(v.Kernel())
		return nil
	}
)

type App struct {
	kernel   *cli.App
	config   Config
	flags    []cli.Flag
	commands []*command.Command
	action   Action
}

func NewApp(options ...Option) *App {
	return NewAppWithConfig(newConfig(options...))
}

func NewAppWithConfig(config Config) *App {
	return &App{
		kernel:   handleAppConfig(cli.NewApp(), config),
		config:   config,
		flags:    make([]cli.Flag, 0),
		commands: make([]*command.Command, 0),
		action:   nil,
	}
}

func (app *App) Kernel() *cli.App {
	// handle if it is a directly created structure
	// 处理如果是直接创建的结构体
	if app.kernel == nil {
		app.kernel = cli.NewApp()
		app.flags = make([]cli.Flag, 0)
		app.commands = make([]*command.Command, 0)
	}
	return app.kernel
}

type Value = flag.Value
type Action func(v Value) error

func (app *App) WithAction(action Action) *App {
	if action != nil {
		app.action = action
		app.Kernel().Action = func(c *cli.Context) error {
			return action(flag.NewValue(c))
		}
	}
	return app
}

func (app *App) Run() error {
	sort.Sort(cli.FlagsByName(app.flags))
	app.Kernel().Flags = app.flags
	for _, cmd := range app.commands {
		app.Kernel().Commands = append(app.Kernel().Commands, *cmd.Kernel())
	}
	//if app.action == nil {
	//	app.Kernel().Action = nilAction
	//}
	return app.Kernel().Run(os.Args)
}

func (app *App) OnBeforeAction(action Action) *App {
	if action != nil {
		app.Kernel().Before = func(c *cli.Context) error {
			return action(flag.NewValue(c))
		}
	}
	return app
}

func (app *App) OnAfterAction(action Action) *App {
	if action != nil {
		app.Kernel().After = func(c *cli.Context) error {
			return action(flag.NewValue(c))
		}
	}
	return app
}

func (app *App) AddCommand(command *command.Command) *App {
	app.commands = append(app.commands, command)
	return app
}

func (app *App) Commands() []*command.Command {
	return app.commands
}

func (app *App) Action() Action {
	return app.action
}

func (app *App) Flags() []cli.Flag {
	return app.flags
}

type CommandNotFoundHandler = func(v Value, s string)

// WithCommandNotFoundHandler Execute this function if the proper command cannot be found
func (app *App) WithCommandNotFoundHandler(handler CommandNotFoundHandler) *App {
	if handler != nil {
		app.Kernel().CommandNotFound = func(context *cli.Context, s string) {
			handler(flag.NewValue(context), s)
		}
	}
	return app
}

type UsageErrorHandler = func(v Value, err error, isSubcommand bool) error

// WithUsageErrorHandler Execute this function if a usage error occurs
func (app *App) WithUsageErrorHandler(handler UsageErrorHandler) *App {
	if handler != nil {
		app.Kernel().OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
			return handler(flag.NewValue(context), err, isSubcommand)
		}
	}
	return app
}

type ExitErrHandlerFunc = func(v Value, err error)

// WithExitErrHandler Execute this function to handle ExitErrors.
func (app *App) WithExitErrHandler(handler ExitErrHandlerFunc) *App {
	if handler != nil {
		app.Kernel().ExitErrHandler = func(context *cli.Context, err error) {
			handler(flag.NewValue(context), err)
		}
	}
	return app
}

func (app *App) WithFlagInt(name string, value int, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.IntFlag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

func (app *App) WithFlagInt64(name string, value int64, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.Int64Flag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

func (app *App) WithFlagUint(name string, value uint, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.UintFlag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

func (app *App) WithFlagUint64(name string, value uint64, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.Uint64Flag{
		Name: name, Value: value, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

// WithFlagIntSlice example: main.ext -names Bob -names Tom -names Lisa
func (app *App) WithFlagIntSlice(name string, value []int, usage string, required bool, hidden ...bool) *App {
	intFlag := cli.IntSliceFlag{
		Name: name, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	}
	*intFlag.Value = append(*intFlag.Value, value...)
	app.flags = append(app.flags, intFlag)
	return app
}

func (app *App) WithFlagString(name string, value string, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.StringFlag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

func (app *App) WithFlagStringSlice(name string, value []string, usage string, required bool, hidden ...bool) *App {
	stringFlag := cli.StringSliceFlag{
		Name: name, Usage: usage,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	}
	*stringFlag.Value = append(*stringFlag.Value, value...)
	app.flags = append(app.flags, stringFlag)
	return app
}

func (app *App) WithFlagBool(name string, value bool, usage string, required bool, hidden ...bool) *App {
	if value {
		app.flags = append(app.flags, cli.BoolTFlag{
			Name: name, Usage: usage, Required: required, Hidden: len(hidden) > 0 && hidden[0],
		})
	} else {
		app.flags = append(app.flags, cli.BoolFlag{
			Name: name, Usage: usage, Required: required, Hidden: len(hidden) > 0 && hidden[0],
		})
	}
	return app
}

func (app *App) WithFlagFloat64(name string, value float64, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.Float64Flag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}

func (app *App) WithFlagDuration(name string, value time.Duration, usage string, required bool, hidden ...bool) *App {
	app.flags = append(app.flags, cli.DurationFlag{
		Name: name, Usage: usage, Value: value,
		Required: required, Hidden: len(hidden) > 0 && hidden[0],
	})
	return app
}
