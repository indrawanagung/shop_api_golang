package service

import (
	"github.com/indrawanagung/shop_api_golang/model/web"
)

type AddressServiceInterface interface {
	FindAllByUserID(userID string) []web.AddressResponse
	FindByID(addressID string) web.AddressResponse
	Save(addressRequest web.AddressCreateOrUpdateRequest) string
	Update(addressRequest web.AddressCreateOrUpdateRequest, addressID string)
	Delete(addressID string)
}
