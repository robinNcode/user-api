package main

import (
	"github.com/gin-gonic/gin"
	"go-api/config/database"
	"go-api/internal/models"
	"go-api/internal/routes"
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