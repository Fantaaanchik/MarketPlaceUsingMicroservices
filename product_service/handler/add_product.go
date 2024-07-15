package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/product_service/models"
	"net/http"
)

func (h Handler) AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	err := h.Service.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot add product"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}
