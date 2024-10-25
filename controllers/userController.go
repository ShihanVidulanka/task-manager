// controllers/userController.go
package controllers

import (
	"log"
	"net/http"
	"task-manager/models"
	"task-manager/services"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users.
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := services.GetAllUsers(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser retrieves a user by ID.
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles user registration.
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v\n", err) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create user using the service layer
	if err := services.CreateUser(&user); err != nil {
		log.Printf("Error creating user: %v\n", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles user updates.
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateUser(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles user deletion.
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
