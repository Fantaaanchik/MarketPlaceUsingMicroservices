package handler

import (
	"github.com/gin-gonic/gin"
	"marketplace/user_service/models"
	"marketplace/user_service/service"
)

type ServiceInterface interface {
	GetUser(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	CreateProduct(user models.User) error
}

type Handler struct {
	Service *service.Service
	Engine  *gin.Engine
}

func New(service *service.Service, engine *gin.Engine) *Handler {
	return &Handler{
		Service: service,
		Engine:  engine,
	}
}
