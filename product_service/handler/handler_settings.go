package handler

import (
	"github.com/gin-gonic/gin"
	"marketplace/product_service/models"
	"marketplace/product_service/service"
)

type ServiceInterface interface {
	GetProduct(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(user models.Product) error
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
