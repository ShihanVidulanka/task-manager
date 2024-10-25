// controllers/taskController.go
package controllers

import (
	"net/http"
	"task-manager/models"
	"task-manager/services"

	"github.com/gin-gonic/gin"
)

// GetTasks retrieves all tasks.
func GetTasks(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by ID.
func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask handles task creation.
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, exists := c.Get("userID") // Get UserID from context/session
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	task.UserID = userID.(uint) // Type assertion to uint

	if err := services.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// UpdateTask handles task updates.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateTask(id, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DeleteTask handles task deletion.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete task"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
