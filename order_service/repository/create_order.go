package repository

import (
	"fmt"
	"marketplace/order_service/models"
)

func (r *Repository) CreateOrder(user models.Order) error {

	if err := r.db.Create(&user).Error; err != nil {
		wrappedError := fmt.Errorf("cannot create order, Error description: %e", err.Error())
		return wrappedError
	}
	return nil
}
