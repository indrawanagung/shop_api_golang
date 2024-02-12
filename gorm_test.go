package main

import (
	"fmt"
	"github.com/indrawanagung/shop_api_golang/db"
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"testing"
)

//func OpenConnection() *gorm.DB {
//	dsn := "host=localhost user=root password=secret dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic(err)
//	}
//	return db
//}

var connectionDB = db.OpenConnection()

func TestGetCities(t *testing.T) {
	var cities []domain.City

	err := connectionDB.Find(&cities).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(cities)
}

func TestGetAdress(t *testing.T) {
	var address domain.Address

	err := connectionDB.Model(&domain.Address{}).Joins("City").First(&address).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(address.User)
	fmt.Println(address.City.City)
}

func TestGetUsers(t *testing.T) {
	var users domain.User

	err := connectionDB.Model(&domain.User{}).Preload("Addresses.City").First(&users).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(users.Addresses[0].City.City)
}

func TestGetStores(t *testing.T) {
	var stores []domain.Store

	err := connectionDB.Model(&domain.Store{}).Joins("User").Find(&stores).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(stores[0].User.FullName)
}

func TestGetWarehouses(t *testing.T) {
	var warehouse domain.Warehouses

	err := connectionDB.Model(&domain.Warehouses{}).Joins("Status.CategoryStatus").Joins("Address.City").Find(&warehouse).Error
	if err != nil {
		panic(err)
	}
	//fmt.Println(warehouse.Status.CategoryStatus)
	fmt.Println(warehouse.Address.City.City)
}

func TestGetCategoryStatus(t *testing.T) {
	var categoryStatus domain.CategoryStatus

	err := connectionDB.Model(&domain.CategoryStatus{}).Find(&categoryStatus).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(categoryStatus)
}

func TestGetStatus(t *testing.T) {
	var status domain.Status

	err := connectionDB.Model(&domain.Status{}).Joins("CategoryStatus").Find(&status).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(status.CategoryStatus)
}

func TestGetProductCategory(t *testing.T) {
	var productCategories []domain.ProductCategory

	err := connectionDB.Model(&domain.ProductCategory{}).Find(&productCategories).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(productCategories)
}
