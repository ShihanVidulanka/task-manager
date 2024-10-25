// middleware/rateLimitMiddleware.go
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var rateLimiter = time.NewTicker(1 * time.Second) // Simple rate limiter for 1 request per second

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case <-rateLimiter.C:
			c.Next()
		default:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
		}
	}
}
