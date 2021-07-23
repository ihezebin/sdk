package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk-go/httpserver/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Name() (name string)
	AddMiddlewares(middlewares ...middleware.Middleware) Server
	GetEngine() (engine *gin.Engine)
	Run(ctx context.Context) error
	Close(ctx context.Context) error
	OnShutdown(f func()) Server
	BeforeRun(f func()) Server
	Route(routes Router) Server
}

type server struct {
	http.Server
	config     Config
	beforeRun  []func()
	onShutdown []func()
}

func NewServer(confs ...ConfigFunc) Server {
	return NewServerWithConfig(newConfig(confs...))
}

func NewServerWithConfig(config Config) Server {
	s := &server{config: config}
	gin.SetMode(config.Mode)
	engine := gin.New()
	// default Use middleware
	engine.Use()
	// user set middleware
	engine.Use(config.Middlewares...)
	s.Handler = engine
	s.Addr = fmt.Sprintf(":%d", config.Port)
	return s
}

func (s *server) Route(routes Router) Server {
	routes(s.GetEngine())
	return s
}

func (s *server) Name() string {
	return s.config.Name
}

func (s *server) Run(ctx context.Context) error {
	// register func before run
	for _, f := range s.beforeRun {
		f()
	}
	// register func on shutdown
	for _, f := range s.onShutdown {
		s.RegisterOnShutdown(f)
	}
	// server listenAndServe
	ch := make(chan os.Signal)
	go func() {
		log.Printf("http server is starting in port:%d", s.config.Port)
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Printf("http server ListenAndServe err:%v\n", err)
				ch <- sigerr
			}
			log.Println("http server closed")
		}
	}()
	// handle signal, to elegant closing server
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, sigerr)
	log.Printf("got signal %v, http server exit\n", <-ch)
	return s.Close(ctx)
}

func (s *server) OnShutdown(f func()) Server {
	s.onShutdown = append(s.onShutdown, f)
	return s
}

func (s *server) BeforeRun(f func()) Server {
	s.beforeRun = append(s.beforeRun, f)
	return s
}

func (s *server) Close(ctx context.Context) error {
	log.Println("http server is closing...")
	return s.Shutdown(ctx)
}

func (s *server) AddMiddlewares(middlewares ...middleware.Middleware) Server {
	s.GetEngine().Use(middlewares...)
	return s
}

func (s *server) GetEngine() *gin.Engine {
	return s.Handler.(*gin.Engine)
}

type signalErr string

const sigerr = signalErr("httpserver err signal")

func (s signalErr) Signal() {

}
func (s signalErr) String() string {
	return "listen and serve err"
}
