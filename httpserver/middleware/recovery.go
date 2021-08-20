package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whereabouts/sdk/logger"
	"net/http"
	"runtime/debug"
)

// Recovery global exception handling middleware
func Recovery() Middleware {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Print error and stack information
				logger.Errorf("panic: %v", err)
				logger.Errorf("stack: %s", debug.Stack())
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
