package main

import (
	"github.com/gin-gonic/gin"
	"user-api/internal/config"
	"user-api/internal/models"
	"user-api/internal/routes"
)

func main() {
	// DB connect
	database.Connect()

	// Auto migrate (like Laravel migrations quick version)
	database.DB.AutoMigrate(&models.User{})

	// Router
	r := gin.Default()

	// Routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":8082")
}