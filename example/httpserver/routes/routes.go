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
	httpserver.Route(engine, http.MethodGet, "/standard", handlers.HelloStandardHandler)
	httpserver.Route(engine, http.MethodPost, "/standard", handlers.HelloStandardHandler)
	httpserver.Route(engine, http.MethodPost, "/file", handlers.HelloFileHandler)
	httpserver.Route(engine, http.MethodPost, "/multiple", handlers.HelloMultipleFilesHandler)
	// child route
	hello := engine.Group("hello")
	{
		httpserver.Route(hello, http.MethodGet, "/standard", handlers.HelloStandardHandler)
		httpserver.Route(hello, http.MethodPost, "/file", handlers.HelloFileHandler)
		httpserver.Route(hello, http.MethodPost, "/multiple", handlers.HelloMultipleFilesHandler)
	}

}
