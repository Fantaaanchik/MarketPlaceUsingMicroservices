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

	r.POST("/create_order", h.CreateOrder)
	r.GET("/get_orders", h.GetAllOrders)
	r.GET("/get_order", h.GetOrder)

	err := r.Run(viper.GetString("service.address"))
	if err != nil {
		log.Fatalf("cannot run order service routes, error: %v", err.Error())
	}

}
