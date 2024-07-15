package db

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"marketplace/user_service/models"
)

func InitDB() *gorm.DB {
	var err error

	dsn := viper.GetString("database.url")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed to connect to DB, %v", err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logrus.Println("cannot migrate model User to DB, ", err.Error())
	}

	return db
}
