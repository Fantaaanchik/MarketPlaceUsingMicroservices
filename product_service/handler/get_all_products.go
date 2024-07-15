package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h Handler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		logrus.Println("", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!"})
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
