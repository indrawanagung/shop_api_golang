package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/exception"
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
	err, cities := c.CityRepository.FindAll(c.DB)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	return cities
}

func (c CityServiceImpl) FindByID(cityID string) domain.City {
	err, city := c.CityRepository.FindByID(c.DB, cityID)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("city id %s is not found", cityID)))
	}
	return city
}
