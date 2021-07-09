package server

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func PreMiddleware(middlewares ...gin.HandlerFunc) {
	gServer.GetEngine().Use(middlewares...)
}

func AfterMiddleware(middlewares ...gin.HandlerFunc) {
	gServer.GetEngine().Use(getAfter(middlewares...)...)
}

func (gr *GroupRouter) PreMiddleware(middlewares ...gin.HandlerFunc) {
	gr.Use(middlewares...)
}

func (gr *GroupRouter) AfterMiddleware(middlewares ...gin.HandlerFunc) {
	gr.Use(getAfter(middlewares...)...)
}

func getAfter(middlewares ...gin.HandlerFunc) (hfs []gin.HandlerFunc) {
	hfs = make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		middleware := func(c *gin.Context) {
			c.Next()
			mwV := reflect.ValueOf(middleware)
			if mwV.IsValid() {
				mwV.Call([]reflect.Value{reflect.ValueOf(c)})
			}
		}
		hfs = append(hfs, middleware)
	}
	return
}
