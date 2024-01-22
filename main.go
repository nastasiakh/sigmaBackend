package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sigma/internal/handlers/di"
	"sigma/internal/handlers/middlewars"
	"sigma/internal/handlers/routes"
)

func main() {
	router := gin.Default()
	router.Use(middlewars.CorsMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	container := di.BuildContainer()

	routes.AddRoutes(router, container)

	router.Run(":8080")
}
