// services/taskService.go
package services

import (
	"task-manager/config"
	"task-manager/models"
)

// CreateTask handles the logic for creating a task.
func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

// GetAllTasks retrieves all tasks.
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := config.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID retrieves a task by ID.
func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates the task details.
func UpdateTask(id string, task *models.Task) error {
	return config.DB.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error
}

// DeleteTask deletes a task by ID.
func DeleteTask(id string) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
