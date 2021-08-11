package main

import (
	"github.com/whereabouts/sdk/config"
	"github.com/whereabouts/sdk/logger"
)

func main() {
	err := config.LoadWithCli("c", "./config.json", GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}
