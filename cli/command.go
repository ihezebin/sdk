package cli

import "github.com/urfave/cli"

type Command struct {
	kernel cli.Command
}

func NewCommand() *Command {
	return &Command{}
}

func (command *Command) Kernel() cli.Command {
	return command.kernel
}
