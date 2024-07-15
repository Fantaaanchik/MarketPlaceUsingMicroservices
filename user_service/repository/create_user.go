package repository

import (
	"fmt"
	"marketplace/user_service/models"
)

func (r *Repository) CreateUser(user models.User) error {

	if err := r.db.Create(&user).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get all users from db, Error description: %e", err.Error())
		return wrappedError
	}
	return nil
}
