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
	Close(ctx context.Context)
	OnShutdown(f func()) Server
	BeforeRun(f func()) Server
	Route(routes Router) Server
}

type server struct {
	http.Server
	option     Option
	beforeRun  []func()
	onShutdown []func()
}

func NewServer(opts ...OptionFunc) Server {
	return NewServerWithOption(newOption(opts...))
}

func NewServerWithOption(option Option) Server {
	s := &server{option: option}
	gin.SetMode(option.Mode)
	engine := gin.New()
	// default Use middleware
	engine.Use()
	// user set middleware
	engine.Use(option.Middlewares...)
	s.Handler = engine
	s.Addr = fmt.Sprintf(":%d", option.Port)
	return s
}

func (s *server) Route(routes Router) Server {
	routes(s.GetEngine())
	return s
}

func (s *server) Name() string {
	return s.option.Name
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
	go func() {
		log.Printf("http server:%s is starting in port:%d", s.option.Name, s.option.Port)
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Printf("httpServer ListenAndServe err:%v\n", err)
			}
			log.Println("http server is closed")
		}
	}()
	// handle signal, to elegant closing server
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	log.Printf("got signal %v, httpServer exit\n", <-ch)
	s.Close(ctx)
	return nil
}

func (s *server) OnShutdown(f func()) Server {
	s.onShutdown = append(s.onShutdown, f)
	return s
}

func (s *server) BeforeRun(f func()) Server {
	s.beforeRun = append(s.beforeRun, f)
	return s
}

func (s *server) Close(ctx context.Context) {
	s.Shutdown(ctx)
}

func (s *server) AddMiddlewares(middlewares ...middleware.Middleware) Server {
	s.GetEngine().Use(middlewares...)
	return s
}

func (s *server) GetEngine() *gin.Engine {
	return s.Handler.(*gin.Engine)
}
