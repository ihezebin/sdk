package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery global exception handling middleware
func Recovery() Middleware {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Print error stack information
				log.Printf("panic: %v\n", err)
				debug.PrintStack()
				// Terminate subsequent interface calls
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func RecoveryCustom(recover Middleware) Middleware {
	return createRecoveryCustom(recover)
}

func createRecoveryCustom(custom Middleware) Middleware {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// handle your custom
				custom(c)
			}
		}()
		c.Next()
	}
}
