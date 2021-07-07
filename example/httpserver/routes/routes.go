package routes

import (
	"github.com/whereabouts/web-template/engine/server"
	"github.com/whereabouts/web-template/handlers"
	"github.com/whereabouts/web-template/middleware"
	"net/http"
)

func Routes() {
	// middleware use before route
	server.PreMiddleware(middleware.HelloPreMiddleware)
	server.AfterMiddleware(middleware.HelloAfterMiddleware)
	server.Route(http.MethodGet, "/sayHello", handlers.SayHello)
	server.Route(http.MethodPost, "/fileHello", handlers.FileHello)
	server.Route(http.MethodPost, "/filesHello", handlers.FilesHello)
	// add the child route
	v1 := server.Group("v1")
	{
		v1.Route(http.MethodGet, "/sayHello", handlers.SayHello)
		v1.Route(http.MethodPost, "/fileHello", handlers.FileHello)
		v1.Route(http.MethodPost, "/filesHello", handlers.FilesHello)
	}

}
