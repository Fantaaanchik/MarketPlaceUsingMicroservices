package service

import (
	"marketplace/user_service/models"
)

func (s *Service) CreateProduct(user models.User) error {
	return s.Repository.CreateUser(user)
}
