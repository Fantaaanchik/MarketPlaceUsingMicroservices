package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		logrus.Println("", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!"})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
