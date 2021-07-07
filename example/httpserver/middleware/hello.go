package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloPreMiddleware(context *gin.Context) {
	fmt.Println("hello pre")
}

func HelloAfterMiddleware(context *gin.Context) {
	fmt.Println("hello after")
}
