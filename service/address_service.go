package service

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/model/web"
)

type AddressServiceInterface interface {
	FindAllByUserID(userID string) []domain.Address
	FindByID(addressID string) domain.Address
	Save(addressRequest web.AddressCreateOrUpdateRequest) string
	Update(addressRequest web.AddressCreateOrUpdateRequest, addressID string)
	Delete(addressID string)
}
