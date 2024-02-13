package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type AddressRepositoryInterface interface {
	FindAllByUserID(db *gorm.DB, userID string) []domain.Address
	FindByID(db *gorm.DB, addressID string) (error, domain.Address)
	SaveOrUpdate(db *gorm.DB, address domain.Address)
	Delete(db *gorm.DB, addressID string)
}
