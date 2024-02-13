package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type CityRepositoryInterface interface {
	FindAll(db *gorm.DB) (error, []domain.City)
	FindByID(db *gorm.DB, ID string) (error, domain.City)
}
