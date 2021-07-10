package config

import (
	"fmt"
	configure "github.com/whereabouts/sdk-go/configure"
	"log"
	"testing"
)

func TestLoad(t *testing.T) {
	err := configure.LoadJSON("./application.json", &gConfig)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(gConfig.Env, gConfig.Port)
}
