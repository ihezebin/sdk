package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk/example/httpserver/handlers"
	"github.com/whereabouts/sdk/example/httpserver/middleware"
	"github.com/whereabouts/sdk/httpserver"
	"net/http"
)

func Routes(engine *gin.Engine) {
	// use middleware of before and after
	httpserver.Before(engine, middleware.HelloBeforeMiddleware)
	httpserver.After(engine, middleware.HelloAfterMiddleware)
	// route
	httpserver.Route(engine, http.MethodGet, "/standard", handlers.StandardHandler)
	httpserver.Route(engine, http.MethodPost, "/standard", handlers.StandardHandler)
	httpserver.Route(engine, http.MethodPost, "/file", handlers.StandardFileHandler)
	httpserver.Route(engine, http.MethodPost, "/multiple", handlers.StandardMultipleFilesHandler)
	// child route
	v1 := engine.Group("v1")
	{
		httpserver.Handle(v1, http.MethodGet, "/hello", handlers.HelloHandler)
		httpserver.Handle(v1, http.MethodPost, "/hello", handlers.HelloHandler)
		httpserver.Handle(v1, http.MethodPost, "/file", handlers.HelloFileHandler)
		httpserver.Handle(v1, http.MethodPost, "/multiple", handlers.HelloMultipleFilesHandler)
	}

}
