package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h Handler) GetUser(c *gin.Context) {
	ids := c.Query("id")

	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong id type, need number"})
		return
	}

	user, err := h.Service.GetUser(id)
	if err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!"})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
