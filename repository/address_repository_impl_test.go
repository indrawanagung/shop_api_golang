package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var addressRepository = NewAddressRepository()
var address = domain.Address{
	ID:         util.GenerateUUID(),
	Name:       "alamat 1",
	CityID:     "1",
	PostalCode: "9000",
	UserID:     "1",
	Timestamp: domain.Timestamp{
		CreatedAt: util.GetUnixTimestamp(),
	},
}

func SaveAddress(db *gorm.DB, address domain.Address) {
	addressRepository.SaveOrUpdate(db, address)
}

func TestAddressRepositoryImpl_FindAllByUserID(t *testing.T) {
	SaveAddress(dbTest, address)
	addresses := addressRepository.FindAllByUserID(dbTest, address.UserID)
	assert.NotEqual(t, 0, len(addresses))
	addressRepository.Delete(dbTest, address.ID)
}

func TestAddressRepositoryImpl_FindByID(t *testing.T) {
	SaveAddress(dbTest, address)
	err, addressResponse := addressRepository.FindByID(dbTest, address.ID)
	assert.Nil(t, err)
	assert.Equal(t, address.ID, addressResponse.ID)
	assert.Equal(t, address.User, addressResponse.User)
	assert.Equal(t, address.CityID, addressResponse.CityID)

	addressRepository.Delete(dbTest, address.ID)

	//error address is not found
	err, _ = addressRepository.FindByID(dbTest, "987987")
	assert.NotNil(t, err)
}

func TestAddressRepositoryImpl_Delete(t *testing.T) {
	SaveAddress(dbTest, address)
	addressRepository.Delete(dbTest, address.ID)
}
