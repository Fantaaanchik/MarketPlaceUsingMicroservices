package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/order_service/models"
	"net/http"
)

func (h Handler) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	err := h.Service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create order"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
