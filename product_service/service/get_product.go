package service

import "marketplace/product_service/models"

func (s *Service) GetProduct(id int) (models.Product, error) {
	return s.Repository.GetProduct(id)
}
