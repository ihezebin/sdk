package httpserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Name() (name string)
	Run(ctx context.Context) (err error)
	AddMiddlewares(middlewares ...Middleware)
	GetEngine() (engine *gin.Engine)
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
	// default Use middles
	// 默认中间件， FlowControlTag，错误恢复
	engine.Use()
	engine.Use(option.Middlewares...) // user set middles
	s.Handler = engine
	return s
}

func (s *server) Name() string {
	return s.option.Name
}

func (s *server) Run(ctx context.Context) error {

	//for _, f := range s.onShutdown { // 注册用户自定义的关闭函数
	//	srv.RegisterOnShutdown(f)
	//}
	//
	//// check if inside docker, for local debug.
	//if outsideContainer() { // 如果检测出运行环境不是在容器内部，则替换request.ctx为context.Background()，方便在本地调试
	//	middles.MutateRequest = changeRequest
	//}
	//
	//// server listenAndServe
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		if !errors.Is(err, http.ErrServerClosed) {
	//			logger.Errorf(ctx, "httpServer ListenAndServe err:%v", err)
	//		}
	//		logger.Infof(ctx, "http server is closed")
	//	}
	//}()

	// handle signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	log.Printf("got signal %v, exit", <-ch)
	s.Shutdown(ctx)
	return nil
}

func (s *server) RegisterOnShutdown(f func()) {
	s.RegisterOnShutdown(f)
}

func (s *server) AddMiddlewares(middlewares ...Middleware) {
	s.GetEngine().Use(middlewares...)
}

func (s *server) GetEngine() *gin.Engine {
	return s.Handler.(*gin.Engine)
}
