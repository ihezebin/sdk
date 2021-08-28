package main

import (
	"errors"
	"fmt"
	"github.com/whereabouts/sdk/cli"
	"github.com/whereabouts/sdk/cli/command"
	"github.com/whereabouts/sdk/cli/flag"
	"github.com/whereabouts/sdk/logger"
)

func main() {
	// sub command
	addCmd := command.NewCommand(
		command.WithName("add"),
		command.WithDescription("Add file contents to the index"),
	)
	addCmd = addCmd.WithAction(func(v flag.Value) error {
		fmt.Println("succeed to add these files:", v.Args())
		fmt.Println("config:", v.String("config"))
		fmt.Println("global config:", v.GlobalString("config"))
		return nil
	})

	// app
	gitApp := cli.NewApp(
		cli.WithName("git"),
		cli.WithAuthor("Korbin"),
		cli.WithDescription("this is a git app"),
		cli.WithVersion("v1.0.0"),
	)
	gitApp = gitApp.WithFlagString("config, c", "./config.json", "config file path", false)
	gitApp = gitApp.WithFlagBool("bool, b", true, "test bool", false)
	gitApp = gitApp.AddCommand(addCmd)
	gitApp = gitApp.OnBeforeAction(func(v cli.Value) error {
		fmt.Println("git before")
		return nil
	})
	gitApp = gitApp.OnAfterAction(func(v cli.Value) error {
		fmt.Println("git after")
		return nil
	})
	gitApp = gitApp.WithAction(func(v cli.Value) error {
		fmt.Println(v.String("config"))
		fmt.Println(v.Bool("b"))
		return errors.New("default error happened")
	})
	gitApp = gitApp.WithExitErrHandler(func(v cli.Value, err error) {
		fmt.Println("WithExitErrHandler:", err)
	})
	if err := gitApp.Run(); err != nil {
		logger.Error(err.Error())
	}
}
