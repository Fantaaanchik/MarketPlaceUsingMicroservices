package repository

import (
	"fmt"
	"marketplace/user_service/models"
)

func (r *Repository) GetUser(id int) (models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get user with this id, error: %s", err.Error())
		return models.User{}, wrappedError
	}

	return user, nil
}
