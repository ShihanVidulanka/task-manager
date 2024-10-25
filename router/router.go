package router

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/logout", controllers.LogoutUser)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.SessionMiddleware()) // Apply session middleware here

	taskGroup := authorized.Group("/tasks")
	{
		taskGroup.GET("", controllers.GetTasks)
		taskGroup.POST("", controllers.CreateTask)
		taskGroup.GET("/:id", controllers.GetTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}

	userGroup := authorized.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
