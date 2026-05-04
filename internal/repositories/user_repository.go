package repositories

import (
	"go-api/internal/database"
	"go-api/internal/models"
)

type UserRepository struct{}

func (r UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r UserRepository) Create(user models.User) (models.User, error) {
	err := database.DB.Create(&user).Error
	return user, err
}