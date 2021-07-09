package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/example/httpserver/handlers"
	"github.com/whereabouts/sdk-go/httpserver"
	"net/http"
)

type Req struct {
}

type Rsp struct {
}

func main() {
	httpserver.NewServer(
		httpserver.WithName("app"),
		httpserver.WithPort(8080),
		httpserver.WithMode(httpserver.ModeTest),
		httpserver.WithMiddles(),
	).Route(func(engine *gin.Engine) {
		httpserver.Route(engine, http.MethodGet, "ping", handlers.SayHello)
	}).BeforeRun(func() {
		// init db
	}).OnShutdown(func() {
		// close db
	}).Run(context.Background())
}
