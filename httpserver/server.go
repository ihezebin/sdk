package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/httpserver/hook"
	"github.com/whereabouts/sdk/httpserver/middleware"
	"github.com/whereabouts/sdk/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Name() (name string)
	AddMiddlewares(middlewares ...middleware.Middleware) Server
	Kernel() (engine *gin.Engine)
	Run(ctx context.Context) error
	Close(ctx context.Context) error
	OnShutdown(shutdownHook hook.ShutdownHook) Server
	OnBeforeRun(hook.RunHook) Server
}

type server struct {
	http.Server
	config      Config
	onBeforeRun []hook.RunHook
	onShutdown  []hook.ShutdownHook
}

func NewServer(options ...Option) Server {
	return NewServerWithConfig(newConfig(options...))
}

func NewServerWithConfig(config Config) Server {
	s := &server{config: config}
	gin.SetMode(config.Mode)
	engine := gin.New()
	// default Use middleware
	engine.Use()
	// user set middleware
	engine.Use(config.middlewares...)
	s.Handler = engine
	s.Addr = fmt.Sprintf(":%d", config.Port)
	return s
}

func (s *server) Name() string {
	return s.config.Name
}

func (s *server) Run(ctx context.Context) error {
	// register func before run
	for _, beforeRun := range s.onBeforeRun {
		beforeRun()
	}
	// register func on shutdown
	for _, shutdown := range s.onShutdown {
		s.RegisterOnShutdown(shutdown)
	}
	// server listenAndServe
	ch := make(chan os.Signal)
	go func() {
		logger.Infof("http server is starting in port:%d", s.config.Port)
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Printf("http server ListenAndServe err:%v\n", err)
				ch <- sigErr
			}
			logger.Println("http server closed")
		}
	}()
	// handle signal, to elegant closing server
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, sigErr)
	logger.Printf("got signal %v, http server exit\n", <-ch)
	return s.Close(ctx)
}

func (s *server) OnShutdown(f hook.ShutdownHook) Server {
	s.onShutdown = append(s.onShutdown, f)
	return s
}

func (s *server) OnBeforeRun(f hook.RunHook) Server {
	s.onBeforeRun = append(s.onBeforeRun, f)
	return s
}

func (s *server) Close(ctx context.Context) error {
	logger.Println("http server is closing...")
	return s.Shutdown(ctx)
}

func (s *server) AddMiddlewares(middlewares ...middleware.Middleware) Server {
	s.Kernel().Use(middlewares...)
	return s
}

func (s *server) Kernel() *gin.Engine {
	return s.Handler.(*gin.Engine)
}

type signalErr string

const sigErr = signalErr("httpserver err signal")

func (s signalErr) Signal() {

}
func (s signalErr) String() string {
	return "listen and serve err"
}
