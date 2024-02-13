package util

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/db"
	"gorm.io/gorm"
)

func DBTest(path string) *gorm.DB {
	config, err := LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	conn := db.OpenConnection(config.DBSource)
	return conn
}
