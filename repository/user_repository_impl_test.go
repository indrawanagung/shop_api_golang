package repository

import (
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userRepository = NewUserRepository()
var dbTest = util.DBTest("../")
var user domain.User = domain.User{
	ID:           util.GenerateUUID(),
	FullName:     "Khamzat Chimaev",
	EmailAddress: "khamzat@gmail.com",
	PhoneNumber:  "089123123123",
	Password:     "123",
	Timestamp: domain.Timestamp{
		CreatedAt: util.GetUnixTimestamp(),
	},
	Addresses: nil,
}

func SaveOrUpdateUser(user domain.User) error {
	return userRepository.SaveOrUpdate(dbTest, user)
}

func TestUserRepositoryImpl_FindByID(t *testing.T) {
	//util.TruncateDB(dbTest)

	err := SaveOrUpdateUser(user)
	assert.Nil(t, err)

	err, userResponse := userRepository.FindByID(dbTest, user.ID)
	assert.Equal(t, user.ID, userResponse.ID)
	assert.Equal(t, user.Password, userResponse.Password)
	assert.Equal(t, user.PhoneNumber, userResponse.PhoneNumber)
	assert.Equal(t, user.FullName, userResponse.FullName)
	assert.Equal(t, user.EmailAddress, userResponse.EmailAddress)
	assert.Equal(t, user.Timestamp, userResponse.Timestamp)
}

func TestUserRepositoryImpl_Delete(t *testing.T) {
	err := userRepository.Delete(dbTest, user.ID)
	assert.Nil(t, err)
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	err := SaveOrUpdateUser(user)
	assert.Nil(t, err)

	isEmailExist := userRepository.FindByEmail(dbTest, user.EmailAddress)
	assert.Equal(t, true, isEmailExist)

	err = userRepository.Delete(dbTest, user.ID)
	assert.Nil(t, err)
}
