package repository

import (
	"fmt"
	"marketplace/user_service/models"
)

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get all users, error: ", err.Error())
		return nil, wrappedError
	}
	return users, nil
}
