package service

import "marketplace/order_service/models"

func (s *Service) GetAllOrders() ([]models.Order, error) {
	return s.Repository.GetAllOrders()
}
