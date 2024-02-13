package service

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/repository"
	"gorm.io/gorm"
)

type CityServiceImpl struct {
	CityRepository repository.CityRepositoryInterface
	DB             *gorm.DB
}

func NewCityService(cityRepository repository.CityRepositoryInterface, db *gorm.DB) CityServiceInterface {
	return &CityServiceImpl{
		CityRepository: cityRepository,
		DB:             db,
	}
}

func (c CityServiceImpl) FindAll() []domain.City {
	//TODO implement me
	panic("implement me")
}

func (c CityServiceImpl) FindByID(cityID string) domain.City {
	//TODO implement me
	panic("implement me")
}
