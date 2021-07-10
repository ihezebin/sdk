package main

import (
	"context"
	"github.com/whereabouts/sdk-go/configure"
	"github.com/whereabouts/sdk-go/example/httpserver/config"
	"github.com/whereabouts/sdk-go/example/httpserver/routes"
	"github.com/whereabouts/sdk-go/example/httpserver/server"
	"github.com/whereabouts/sdk-go/httpserver"
	"log"
)

func main() {
	//err := configure.LoadJSON("./config/application.json", config.GetConfig())
	err := configure.LoadJSONWithCmd(config.GetConfig()) // go run main.go -c [config file path]
	if err != nil {
		log.Printf("load config error: %v\n", err)
		return
	}
	ctx := context.Background()
	if err := httpserver.NewServer(
		httpserver.WithName(config.GetConfig().AppName),
		httpserver.WithPort(config.GetConfig().Port),
		httpserver.WithMode(config.GetConfig().Mode),
		httpserver.WithMiddles(),
	).Route(routes.Routes).BeforeRun(server.Init).OnShutdown(server.Close).Run(ctx); err != nil {
		log.Printf("server run with error: %v\n", err)
	}
}
