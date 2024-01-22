package routes

import (
	"github.com/gin-gonic/gin"
	"sigma/internal/handlers/di"
)

func AddRoutes(router *gin.Engine, container *di.Container) {

	userController := container.UserController
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
