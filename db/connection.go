package db

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/Viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	viper := Viper.Config()

	dsn := viper.GetString("DATABASE_CONFIG")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
