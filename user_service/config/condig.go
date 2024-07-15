package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigName("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %v", err.Error())
		return
	}
}
