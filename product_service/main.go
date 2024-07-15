package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/product_service/config"
	"marketplace/product_service/db"
	"marketplace/product_service/handler"
	"marketplace/product_service/repository"
	"marketplace/product_service/service"
)

func main() {
	logrus.Info("Starting Product service ...")

	config.InitConfig()
	dbConn := db.InitDB()
	repositorySettings := repository.New(dbConn)
	serviceSettings := service.New(repositorySettings)

	r := gin.Default()
	handlerSettings := handler.New(serviceSettings, r)

	handlerSettings.AllRoutes()
}
