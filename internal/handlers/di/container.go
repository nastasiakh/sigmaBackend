package di

import (
	"sigma/internal/handlers/controllers"
	"sigma/internal/handlers/repositories"
	"sigma/internal/handlers/services"
)

type Container struct {
	UserController *controllers.UserController
}

func BuildContainer() *Container {
	userRepository := repositories.NewUserRepository()

	userService := services.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
