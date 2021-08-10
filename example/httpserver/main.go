package main

import (
	"context"
	"github.com/whereabouts/sdk/example/httpserver/config"
	"github.com/whereabouts/sdk/example/httpserver/routes"
	"github.com/whereabouts/sdk/example/httpserver/server"
	"github.com/whereabouts/sdk/httpserver"
	"github.com/whereabouts/sdk/httpserver/middleware"
	"github.com/whereabouts/sdk/logger"
)

func main() {
	ctx := context.Background()
	if err := config.Load(); err != nil {
		logger.Fatalf("load config error: %v\n", err)
	}
	if err := httpserver.NewServer(
		httpserver.WithName(config.GetConfig().AppName),
		httpserver.WithPort(config.GetConfig().Port),
		httpserver.WithMode(config.GetConfig().Mode),
		httpserver.WithMiddles(
			middleware.Recovery(),
			middleware.LoggingSimplyRequest(),
			middleware.LoggingSimplyResponse(),
		),
	).Route(routes.Routes).BeforeRun(server.Init).OnShutdown(server.Close).Run(ctx); err != nil {
		logger.Fatalf("server run with error: %v\n", err)
	}
}
