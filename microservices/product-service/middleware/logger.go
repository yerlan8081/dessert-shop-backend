package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[%s] [请求ID: %s] %s %s - %d - 持续时间: %.3fms",
			start.Format(time.RFC3339),
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			status,
			float64(duration.Microseconds())/1000)
	}
}
