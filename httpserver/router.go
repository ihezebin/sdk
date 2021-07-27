package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk/httpserver/middleware"
	"reflect"
)

type Router = func(engine *gin.Engine)

func Route(routes gin.IRoutes, method string, path string, function interface{}) {
	routes.Handle(method, path, middleware.CreateHandlerFunc(function))
}

func Before(routes gin.IRoutes, before ...middleware.Middleware) {
	routes.Use(before...)
}

func After(routes gin.IRoutes, after ...middleware.Middleware) {
	routes.Use(createAfters(after...)...)
}

func createAfters(after ...gin.HandlerFunc) []middleware.Middleware {
	afterSlice := make([]gin.HandlerFunc, 0)
	for _, a := range after {
		a := func(c *gin.Context) {
			c.Next()
			aV := reflect.ValueOf(a)
			if aV.IsValid() {
				aV.Call([]reflect.Value{reflect.ValueOf(c)})
			}
		}
		afterSlice = append(afterSlice, a)
	}
	return afterSlice
}
