package server

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ihezebin/sdk/example/httpserver/handler"
)

func Routes(engine *gin.Engine) {
	new(handlers.HelloHandler).Init(engine)
}
