package repository

import (
	"fmt"
	"marketplace/product_service/models"
)

func (r *Repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get all products, error: %s", err.Error())
		return nil, wrappedError
	}
	return products, nil
}
