package service

import (
	"marketplace/user_service/models"
	"marketplace/user_service/repository"
)

type RepositoryInterface interface {
	GetUser(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	CreateProduct(user models.User) error
}

type Service struct {
	Repository *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
