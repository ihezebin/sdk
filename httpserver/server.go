package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Name() (name string)
	AddMiddlewares(middlewares ...Middleware)
	GetEngine() (engine *gin.Engine)
	Run(ctx context.Context) (err error)
	Shutdown(ctx context.Context)
	RegisterOnShutdown(f func())
}

type server struct {
	http.Server
	option Option
}

func NewServer(opts ...OptionFunc) Server {
	return NewServerWithOption(newOption(opts...))
}

func NewServerWithOption(option Option) Server {
	s := &server{option: option}
	engine := gin.New()
	// default Use middleware
	engine.Use()
	// user set middlewares
	engine.Use(option.Middlewares...)
	s.Handler = engine
	s.Addr = fmt.Sprintf(":%d", option.Port)
	return s
}

func (s *server) Name() string {
	return s.option.Name
}

func (s *server) Run(ctx context.Context) error {

	// server listenAndServe
	go func() {
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
	s.Shutdown(ctx)
	return nil
}

func (s *server) RegisterOnShutdown(f func()) {
	s.RegisterOnShutdown(f)
}

func (s *server) Shutdown(ctx context.Context) {
	s.Shutdown(ctx)
}

func (s *server) AddMiddlewares(middlewares ...Middleware) {
	s.GetEngine().Use(middlewares...)
}

func (s *server) GetEngine() *gin.Engine {
	return s.Handler.(*gin.Engine)
}
