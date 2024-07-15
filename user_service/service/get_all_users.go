package service

import "marketplace/user_service/models"

func (s *Service) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}
