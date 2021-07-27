package main

import (
	"context"
	"github.com/whereabouts/sdk/configure"
	"github.com/whereabouts/sdk/example/httpserver/config"
	"github.com/whereabouts/sdk/example/httpserver/routes"
	"github.com/whereabouts/sdk/example/httpserver/server"
	"github.com/whereabouts/sdk/httpserver"
	"github.com/whereabouts/sdk/httpserver/middleware"
	"log"
)

func main() {
	//err := configure.LoadJSON("./configure/application.json", configure.GetConfig())
	err := configure.LoadJSONWithCmd(config.GetConfig()) // go run main.go -c [configure file path]
	if err != nil {
		log.Printf("load configure error: %v\n", err)
		return
	}
	ctx := context.Background()
	if err := httpserver.NewServer(
		httpserver.WithName(config.GetConfig().AppName),
		httpserver.WithPort(config.GetConfig().Port),
		httpserver.WithMode(config.GetConfig().Mode),
		httpserver.WithMiddles(
			middleware.Recovery(),
			middleware.LoggingRequest(),
			middleware.LoggingResponse(),
		),
	).Route(routes.Routes).BeforeRun(server.Init).OnShutdown(server.Close).Run(ctx); err != nil {
		log.Printf("server run with error: %v\n", err)
	}
}
