package routes

import (
	"github.com/gin-gonic/gin"
	"user-api/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()

	api := r.Group("/api")
	{
		api.GET("/users", userController.Index)
		api.POST("/users", userController.Store)
	}
}