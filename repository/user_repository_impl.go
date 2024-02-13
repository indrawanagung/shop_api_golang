package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) SaveOrUpdate(tx *gorm.DB, user domain.User) error {
	return tx.Save(user).Error
}

func (u UserRepositoryImpl) FindByID(tx *gorm.DB, userID string) (error, domain.User) {
	var user domain.User
	err := tx.Take(&user, "id = ? ", userID).Error
	return err, user
}

func (u UserRepositoryImpl) Delete(tx *gorm.DB, userID string) error {
	user := domain.User{ID: userID}
	return tx.Delete(&user).Error
}

func (u UserRepositoryImpl) FindByEmail(tx *gorm.DB, email string) bool {
	err := tx.Take(&domain.User{}, "email_address = ? ", email).Error
	if err != nil {
		return false
	}
	return true
}
