package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepository() AddressRepositoryInterface {
	return &AddressRepositoryImpl{}
}

func (a AddressRepositoryImpl) FindAllByUserID(db *gorm.DB, userID string) []domain.Address {
	var addresses []domain.Address
	err := db.Where("user_id = ?", userID).Joins("City").Find(&addresses).Error
	fmt.Println(err)
	return addresses
}

func (a AddressRepositoryImpl) FindByID(db *gorm.DB, addressID string) (error, domain.Address) {
	var address domain.Address
	err := db.Joins("City").First(&address, "address.id = ?", addressID).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return errors.New(fmt.Sprintf("address id %s is not found", addressID)), address
	}

	return nil, address
}

func (a AddressRepositoryImpl) SaveOrUpdate(db *gorm.DB, address domain.Address) {
	err := db.Save(&address).Error
	fmt.Println(err)
}

func (a AddressRepositoryImpl) Delete(db *gorm.DB, addressID string) {
	err := db.Delete(&domain.Address{ID: addressID}).Error
	fmt.Println(err)
}
