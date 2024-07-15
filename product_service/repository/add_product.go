package repository

import (
	"fmt"
	"marketplace/product_service/models"
)

func (r *Repository) AddProduct(product models.Product) error {

	if err := r.db.Create(&product).Error; err != nil {
		wrappedError := fmt.Errorf("cannot add new order to db, Error description: %e", err.Error())
		return wrappedError
	}
	return nil
}
