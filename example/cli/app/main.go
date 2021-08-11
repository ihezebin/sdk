package main

import (
	"fmt"
	"github.com/whereabouts/sdk/cli"
	"github.com/whereabouts/sdk/cli/command"
)

func main() {
	command := command.NewCommand(
		command.WithName("do"),
		command.WithCategory("other"),
	).WithFlagString("thing, t", "something", "do some thing", false).
		WithAction(func(v command.Value) error {
			fmt.Println(v.String("t"))
			return nil
		})

	err := cli.NewApp(
		cli.WithAuthor("Korbin"),
		cli.WithDescription("this is a cli app"),
	).
		WithFlagString("config, c", "./config.json", "config file path", false).
		WithFlagBool("bool, b", true, "test bool", false).
		WithCommand(command).
		WithAction(func(v cli.Value) error {
			fmt.Println(v.String("config"))
			fmt.Println(v.Bool("b"))
			return nil
		}).
		Run()
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
