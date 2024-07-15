package repository

import (
	"fmt"
	"marketplace/product_service/models"
)

func (r *Repository) GetProduct(id int) (models.Product, error) {
	var product models.Product

	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get product with this id, error: %s", err.Error())
		return models.Product{}, wrappedError
	}

	return product, nil
}
