package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	SaveOrUpdate(tx *gorm.DB, user domain.User) error
	FindByID(tx *gorm.DB, userID string) (error, domain.User)
	FindByEmail(tx *gorm.DB, email string) bool
	Delete(tx *gorm.DB, userID string) error
}
