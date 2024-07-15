package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/user_service/config"
	"marketplace/user_service/db"
	"marketplace/user_service/handler"
	"marketplace/user_service/repository"
	"marketplace/user_service/service"
)

func main() {
	logrus.Info("Starting User service ...")

	config.InitConfig()
	dbConn := db.InitDB()
	repositorySettings := repository.New(dbConn)
	serviceSettings := service.New(repositorySettings)

	r := gin.Default()
	handlerSettings := handler.New(serviceSettings, r)

	handlerSettings.AllRoutes()
}
