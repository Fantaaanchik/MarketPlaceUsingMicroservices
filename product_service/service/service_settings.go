package service

import (
	"marketplace/product_service/models"
	"marketplace/product_service/repository"
)

type RepositoryInterface interface {
	GetProduct(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(user models.Product) error
}

type Service struct {
	Repository *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
