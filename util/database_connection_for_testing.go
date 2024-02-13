package util

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/db"
	"gorm.io/gorm"
)

func databaseTest() *gorm.DB {
	config, err := LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	conn := db.OpenConnection(config.DBSource)
	return conn
}

var DatabaseTesting = databaseTest()
