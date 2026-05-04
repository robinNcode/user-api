package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{repo: repositories.UserRepository{}}
}

func (s UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s UserService) CreateUser(user models.User) (models.User, error) {
	// you can add validation here (like Laravel FormRequest)
	return s.repo.Create(user)
}