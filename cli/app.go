package cli

import (
	"github.com/urfave/cli"
	"github.com/whereabouts/sdk/cli/command"
	"os"
	"sort"
	"time"
)

var nilAction = func(c *cli.Context) error {
	return nil
}

type App struct {
	kernel   *cli.App
	config   Config
	flags    []cli.Flag
	commands []*command.Command
	action   Action
	help     bool
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
		help:     true,
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

type Value = command.Value
type Action func(v Value) error

func (app *App) WithAction(action Action) *App {
	app.action = action
	app.Kernel().Action = func(c *cli.Context) error {
		return action(command.NewValue(c))
	}
	return app
}

func (app *App) Run() error {
	sort.Sort(cli.FlagsByName(app.flags))
	app.Kernel().Flags = app.flags
	for _, cmd := range app.commands {
		app.Kernel().Commands = append(app.Kernel().Commands, *cmd.Kernel())
	}
	if app.action == nil && !app.help {
		app.Kernel().Action = nilAction
	}
	return app.Kernel().Run(os.Args)
}

func (app *App) Before(action Action) *App {
	app.Kernel().Before = func(c *cli.Context) error {
		return action(command.NewValue(c))
	}
	return app
}

func (app *App) After(action Action) *App {
	app.Kernel().After = func(c *cli.Context) error {
		return action(command.NewValue(c))
	}
	return app
}

func (app *App) WithCommand(command *command.Command) *App {
	app.commands = append(app.commands, command)
	return app
}

func (app *App) WithDefaultHelp(help bool) *App {
	app.help = help
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
