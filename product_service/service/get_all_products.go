package service

import "marketplace/product_service/models"

func (s *Service) GetAllProducts() ([]models.Product, error) {
	return s.Repository.GetAllProducts()
}
