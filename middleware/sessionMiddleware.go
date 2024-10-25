// middleware/sessionMiddleware.go
package middleware

import (
	"net/http"
	"task-manager/config"

	"github.com/gin-gonic/gin"
)

// SessionMiddleware checks if the user is logged in.
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := config.Store.Get(c.Request, "session-name")

		userID, exists := session.Values["userID"]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", userID) // Store userID in context for further use
		c.Next()
	}
}
