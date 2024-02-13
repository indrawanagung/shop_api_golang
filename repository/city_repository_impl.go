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
	//TODO implement me
	panic("implement me")
}

func (c CityRepositoryImpl) FindByID(db *gorm.DB, ID string) (error, domain.City) {
	//TODO implement me
	panic("implement me")
}
