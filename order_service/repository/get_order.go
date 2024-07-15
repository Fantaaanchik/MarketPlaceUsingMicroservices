package repository

import (
	"fmt"
	"marketplace/order_service/models"
)

func (r *Repository) GetOrder(id int) (models.Order, error) {
	var user models.Order

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get order with this id, error: %s", err.Error())
		return models.Order{}, wrappedError
	}

	return user, nil
}
