package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/db"
	"github.com/indrawanagung/shop_api_golang/util"
	"gorm.io/gorm"
)

var conn = DB()

func DB() *gorm.DB {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	conn := db.OpenConnection(config.DBSource)
	return conn
}
