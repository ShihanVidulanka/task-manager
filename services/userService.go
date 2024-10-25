// services/userService.go
package services

import (
	"task-manager/config"
	"task-manager/models"
)

// CreateUser handles user registration logic.
func CreateUser(user *models.User) error {
	// Hash the password before saving (Omitted for simplicity)
	return config.DB.Create(user).Error
}

// GetAllUsers retrieves all users.
func GetAllUsers(users *[]models.User) error {
	return config.DB.Find(users).Error
}

// GetUserByID retrieves a user by ID.
func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates the user details.
func UpdateUser(id string, user *models.User) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

// DeleteUser deletes a user by ID.
func DeleteUser(id string) error {
	return config.DB.Delete(&models.User{}, id).Error
}
