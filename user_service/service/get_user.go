package service

import "marketplace/user_service/models"

func (s *Service) GetUser(id int) (models.User, error) {
	return s.Repository.GetUser(id)
}
