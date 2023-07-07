package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"playground/docs"
	"playground/pkg/handlers"
)

func main() {
	r := gin.Default()

	// API endpoints
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := r.Group("/api/v1")
	{
		api.GET("/users", handlers.ListUsers)
		//api.POST("/users", handlers.AddUser)
		//api.PUT("/users/:id", handlers.UpdateUser)
		//api.DELETE("/users/:id", handlers.DeleteUser)
	}

	// Serve Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
