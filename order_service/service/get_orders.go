package service

import "marketplace/order_service/models"

func (s *Service) GetOrder(id int) (models.Order, error) {
	return s.Repository.GetOrder(id)
}
