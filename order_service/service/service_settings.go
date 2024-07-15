package service

import (
	"marketplace/order_service/models"
	"marketplace/order_service/repository"
)

type RepositoryInterface interface {
	GetOrder(id int) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
	CreateOrder(order models.Order) error
}

type Service struct {
	Repository *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
