// controllers/authController.go
package controllers

import (
	"log"
	"net/http"
	"task-manager/config"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser handles user login.
func LoginUser(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	log.Printf("Error binding JSON: %v\n", foundUser) // Log the error
	if err := config.DB.Where("username = ?", user.Username).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a new session
	session, _ := config.Store.Get(c.Request, "session-name")
	session.Values["userID"] = foundUser.ID
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": foundUser.Username})
}

// LogoutUser handles user logout.
func LogoutUser(c *gin.Context) {
	session, _ := config.Store.Get(c.Request, "session-name")
	// Clear the session
	delete(session.Values, "userID")
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
