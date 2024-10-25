// models/user.go
package models

import "gorm.io/gorm"

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;index" validate:"required,min=3,max=12"`
	Email    string `json:"email" gorm:"unique;not null;index" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=8"`       // exclude from json output store hashed passwords
	Role     string ` json:"role" gorm:"not null;default:user" validate:"required"` // e.g, "admin", "user"
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
