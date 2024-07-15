package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/order_service/config"
	"marketplace/order_service/db"
	"marketplace/order_service/handler"
	"marketplace/order_service/repository"
	"marketplace/order_service/service"
)

func main() {
	logrus.Info("Starting Order service ...")

	config.InitConfig()
	dbConn := db.InitDB()
	repositorySettings := repository.New(dbConn)
	serviceSettings := service.New(repositorySettings)

	r := gin.Default()
	handlerSettings := handler.New(serviceSettings, r)

	handlerSettings.AllRoutes()
}
