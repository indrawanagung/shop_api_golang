package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	city := City{
		ID: "1",
	}

	err := db.Delete(&city).Error
	if err != nil {
		panic(err)
	}
}

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=root password=secret dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

var db = OpenConnection()
