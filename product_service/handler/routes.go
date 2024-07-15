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

	r.POST("/add_product", h.AddProduct)
	r.GET("/get_products", h.GetAllProducts)
	r.GET("/get_product", h.GetProduct)

	err := r.Run(viper.GetString("service.address"))
	if err != nil {
		log.Fatalf("cannot run product service routes, error: %v", err.Error())
	}

}
