package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type CityRepositoryImpl struct {
}

func NewCityRepository() CityRepositoryInterface {
	return &CityRepositoryImpl{}
}

func (c CityRepositoryImpl) FindAll(db *gorm.DB) (error, []domain.City) {
	var cities []domain.City
	err := db.Find(&cities).Error
	return err, cities

}

func (c CityRepositoryImpl) FindByID(db *gorm.DB, ID string) (error, domain.City) {
	var city domain.City
	err := db.Take(&city, "id = ?", ID).Error
	return err, city
}
