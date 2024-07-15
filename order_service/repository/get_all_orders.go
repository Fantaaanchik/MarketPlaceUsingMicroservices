package repository

import (
	"fmt"
	"marketplace/order_service/models"
)

func (r *Repository) GetAllOrders() ([]models.Order, error) {
	var users []models.Order
	if err := r.db.Find(&users).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get all orders, error: %s", err.Error())
		return nil, wrappedError
	}
	return users, nil
}
