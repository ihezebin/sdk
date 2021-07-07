package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk-go/httpserver/configure"
	"github.com/whereabouts/sdk-go/logger"
	"net/http"
)

const ProdMode = "prod"

var gServer *Server

type Server struct {
	http.Server
	config *configure.DefaultConfig
}

func (s *Server) Run() error {
	logger.Infof("server is running in port:%d, env:%s", s.config.Port, s.config.Env)
	return s.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.Shutdown()
}

func (s *Server) SetEngine(engine *gin.Engine) {
	s.Handler = engine
}

func (s *Server) GetEngine() *gin.Engine {
	return s.Handler.(*gin.Engine)
}

func (s *Server) SetAddr(addr string) {
	s.Addr = addr
}

func (s *Server) SetEnv(env string) {
	switch env {
	case ProdMode, gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func NewServer(conf configure.IConfig) *Server {
	// Load the configuration file, you can transfer the parameters of path from the command line like this:
	// "go run main.go -c application.json", default value is "./configure/application.json"
	// another way is that you can name your configure file as "application.json", and put it in the "%project%/configure" directory
	dConf, err := configure.Load(conf)
	if err != nil {
		logger.Fatalf("fail to create server, because of err: %+v", err)
	}
	gServer = createServer(dConf)
	return gServer
}

func DefaultServer() *Server {
	gServer = createServer(configure.DefaultLoad())
	return gServer
}

func createServer(dConf *configure.DefaultConfig) *Server {
	server := &Server{config: dConf}
	// set environment mode，use before gin.New()
	server.SetEnv(server.config.Env)
	// set server engine
	server.SetEngine(gin.New())
	server.SetAddr(fmt.Sprintf(":%d", server.config.Port))
	return server
}

func GetServer() *Server {
	return gServer
}
