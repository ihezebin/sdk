package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloBeforeMiddleware(context *gin.Context) {
	fmt.Println("hello before")
}

func HelloAfterMiddleware(context *gin.Context) {
	fmt.Println("hello after")
}
