package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/httpserver/middleware"
	"reflect"
)

type Router = func(engine *gin.Engine)
type HandlerBefore = gin.HandlerFunc
type HandlerAfter = gin.HandlerFunc

func Route(routes gin.IRoutes, method string, path string, function interface{}) {
	routes.Handle(method, path, middleware.CreateHandlerFunc(function))
}

func RouteBefore(routes gin.IRoutes, befores ...HandlerBefore) {
	routes.Use(befores...)
}

func RouteAfter(routes gin.IRoutes, afters ...HandlerAfter) {
	routes.Use(createAfters(afters...)...)
}

func createAfters(afters ...gin.HandlerFunc) []HandlerAfter {
	afterSlice := make([]gin.HandlerFunc, 0)
	for _, after := range afters {
		after := func(c *gin.Context) {
			c.Next()
			afterV := reflect.ValueOf(after)
			if afterV.IsValid() {
				afterV.Call([]reflect.Value{reflect.ValueOf(c)})
			}
		}
		afters = append(afters, after)
	}
	return afterSlice
}
