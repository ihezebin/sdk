package httpserver

import "github.com/gin-gonic/gin"

type Router = func(engine *gin.Engine)
