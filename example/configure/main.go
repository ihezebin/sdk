package main

import (
	"fmt"
	"github.com/whereabouts/sdk/configure"
	"log"
)

func main() {
	//err := configure.LoadJSONWithCmd(&gConfig)
	err := configure.LoadJSON("./application.json", &gConfig)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(gConfig.Env, gConfig.Port)
}
