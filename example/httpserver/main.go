package main

import (
	"context"
	"github.com/whereabouts/sdk-go/example/httpserver/routes"
	"github.com/whereabouts/sdk-go/example/httpserver/server"
	"github.com/whereabouts/sdk-go/httpserver"
	"log"
)

func main() {
	ctx := context.Background()
	if err := httpserver.NewServer(
		httpserver.WithName("app"),
		httpserver.WithPort(8080),
		httpserver.WithMode(httpserver.ModeDebug),
		httpserver.WithMiddles(),
	).Route(routes.Routes).BeforeRun(server.Init).OnShutdown(server.Close).Run(ctx); err != nil {
		log.Printf("server run with error: %v\n", err)
	}
}
