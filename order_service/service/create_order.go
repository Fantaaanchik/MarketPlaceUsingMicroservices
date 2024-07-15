package service

import (
	"marketplace/order_service/models"
)

func (s *Service) CreateOrder(user models.Order) error {
	return s.Repository.CreateOrder(user)
}
