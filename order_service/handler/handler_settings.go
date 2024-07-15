package handler

import (
	"github.com/gin-gonic/gin"
	"marketplace/order_service/models"
	"marketplace/order_service/service"
)

type ServiceInterface interface {
	GetOrder(id int) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
	CreateOrder(order models.Order) error
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
