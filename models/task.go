// models/task.go
package models

import "gorm.io/gorm"

// Task represents a task in the system.
type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255) not null" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
	UserID      uint   `json:"user_id" gorm:"type:integer not null" validate:"required"`    // Foreign key to User
	Status      string `json:"status" gorm:"type:varchar(20) not null" validate:"required"` // Added Status field
}
