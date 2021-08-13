package main

import (
	"fmt"
	"github.com/whereabouts/sdk/cli"
)

func main() {
	cli.WithFlagString("config, c", "./config", "config file path")
	cli.WithFlagInt("port, p", 8080, "server listen port")
	if value, err := cli.Run(); err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println(value.String("c"))
		fmt.Println(value.String("p"))
	}
}
