package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerBro() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)
		clientIP := c.ClientIP()

		fmt.Printf("%s | %d | %10v | %s | %s %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			status,
			duration,
			clientIP,
			method,
			path,
		)

	}
}
