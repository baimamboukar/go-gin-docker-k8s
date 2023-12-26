package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// The`LoggerMiddleware` function
// --------------------------------------------------------
// returns a Gin `Handler Function` that measures the request time, and Logs basic information about the request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}
