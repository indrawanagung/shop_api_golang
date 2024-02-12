package Viper

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func Config() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath("../..")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
