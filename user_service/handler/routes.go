package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func (h Handler) AllRoutes() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "page not found"})
	})

	r.POST("/create_user", h.CreateUser)
	r.GET("/get_users", h.GetAllUsers)
	r.GET("/get_user", h.GetUser)

	err := r.Run(viper.GetString("service.address"))
	if err != nil {
		log.Fatalf("cannot run user service routes, error: %v", err.Error())
	}

}
