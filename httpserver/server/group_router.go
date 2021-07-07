package server

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/httpserver/hanlder"
)

type GroupRouter struct {
	*gin.RouterGroup
}

func (gr *GroupRouter) Route(method string, path string, function interface{}) {
	gr.Handle(method, path, hanlder.CreateHandlerFunc(function))
}

func (gr *GroupRouter) Group(relativePath string) *GroupRouter {
	return &GroupRouter{gr.RouterGroup.Group(relativePath)}
}

func Group(relativePath string) *GroupRouter {
	return &GroupRouter{gServer.GetEngine().Group(relativePath)}
}
