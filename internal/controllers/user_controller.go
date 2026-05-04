package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"user-api/internal/models"
	"user-api/internal/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController() UserController {
	return UserController{service: services.NewUserService()}
}

func (uc UserController) Index(c *gin.Context) {
	users, err := uc.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc UserController) Store(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := uc.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}