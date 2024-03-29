package service

import (
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var addressRepository = repository.NewAddressRepository()
var addressService = NewAddressService(database, addressRepository, validate)

var addressRequest = web.AddressCreateOrUpdateRequest{
	Name:       "alamat 2",
	CityID:     "1",
	PostalCode: "900",
	UserID:     "1",
}

func SaveAddress() string {
	return addressService.Save(addressRequest)
}

func TestAddressServiceImpl_Delete(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	id := SaveAddress()
	addressService.Delete(id)
	userService.FindByID(id)

}

func TestAddressServiceImpl_FindAllByUserID(t *testing.T) {
	id := SaveAddress()
	addresses := addressService.FindAllByUserID(addressRequest.UserID)
	assert.NotEqual(t, 0, len(addresses))

	addressService.Delete(id)
}

func TestAddressServiceImpl_FindByID(t *testing.T) {
	id := SaveAddress()
	address := addressService.FindByID(id)
	assert.Equal(t, addressRequest.Name, address.Name)
	assert.Equal(t, addressRequest.PostalCode, address.PostalCode)
}

func TestAddressServiceImpl_Update(t *testing.T) {
	id := SaveAddress()

	addressService.FindByID(id)

	addressRequest.PostalCode = "91"
	addressService.Update(addressRequest, id)

	address := addressService.FindByID(id)
	assert.Equal(t, addressRequest.PostalCode, address.PostalCode)

	addressService.Delete(id)
}
