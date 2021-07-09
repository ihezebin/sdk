package server

import (
	"github.com/whereabouts/sdk-go/httpserver/middleware"
)

type Router func()

func (s *Server) Router(r Router) *Server {
	r()
	return s
}

func Route(method string, path string, function interface{}) {
	gServer.GetEngine().Handle(method, path, middleware.CreateHandlerFunc(function))
}

//func setRouter(s *Server, r Router) {
//	r()
//}
