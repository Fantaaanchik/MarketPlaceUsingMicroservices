package service

import (
	"marketplace/product_service/models"
)

func (s *Service) AddProduct(product models.Product) error {
	return s.Repository.AddProduct(product)
}
