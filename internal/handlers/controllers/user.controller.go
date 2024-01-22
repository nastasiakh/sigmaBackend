package controllers

import (
	"github.com/gin-gonic/gin"
	"sigma/internal/entities"
	"sigma/internal/handlers/services"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsersService()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive users"})
	}
	ctx.JSON(200, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	user := c.userService.GetUserByIDService(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive user"})
		return
	}
	ctx.JSON(200, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := c.userService.CreateUserService(user)
	if err != nil {
		return
	}
	ctx.JSON(201, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)

	c.userService.UpdateUserService(user)
	ctx.JSON(200, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	err = c.userService.DeleteUserService(uint(id))
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted"})
}
