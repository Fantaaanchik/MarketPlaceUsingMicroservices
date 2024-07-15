package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/user_service/models"
	"net/http"
)

func (h Handler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	err := h.Service.CreateProduct(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create user"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
