// main.go
package main

import (
	"task-manager/config"
	"task-manager/models"
	"task-manager/router"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.User{}, &models.Task{}) // Run migrations
	r := router.SetupRouter()
	r.Run(":8080")
}
