package main

import (
	"fmt"
	"github.com/whereabouts/sdk/cli"
	"github.com/whereabouts/sdk/cli/command"
	"github.com/whereabouts/sdk/cli/flag"
)

func main() {
	cmd := command.NewCommand(
		command.WithName("do"),
		command.WithCategory("other"),
	).WithFlagString("thing, t", "something", "do some thing", false).
		SetAction(func(v flag.Value) error {
			fmt.Println(v.String("t"))
			return nil
		})

	err := cli.NewApp(
		cli.WithAuthor("Korbin"),
		cli.WithDescription("this is a cli app"),
	).
		WithFlagString("config, c", "./config.json", "config file path", false).
		WithFlagBool("bool, b", true, "test bool", false).
		SetCommand(cmd).
		OnBeforeAction(func(v cli.Value) error {
			fmt.Println("before")
			return nil
		}).OnAfterAction(func(v cli.Value) error {
		fmt.Println("after")
		return nil
	}).
		SetAction(func(v cli.Value) error {
			fmt.Println(v.String("config"))
			fmt.Println(v.Bool("b"))
			return nil
		}).
		Run()
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
