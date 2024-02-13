package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/db"
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var userRepository = NewUserRepository()

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

var conn = DB()

func DB() *gorm.DB {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	conn := db.OpenConnection(config.DBSource)
	return conn
}

func SaveOrUpdateUser(user domain.User) error {
	return userRepository.SaveOrUpdate(DB(), user)
}

func TestUserRepositoryImpl_FindByID(t *testing.T) {
	//util.TruncateDB(conn)

	err := SaveOrUpdateUser(user)
	assert.Nil(t, err)

	err, userResponse := userRepository.FindByID(conn, user.ID)
	assert.Equal(t, user.ID, userResponse.ID)
	assert.Equal(t, user.Password, userResponse.Password)
	assert.Equal(t, user.PhoneNumber, userResponse.PhoneNumber)
	assert.Equal(t, user.FullName, userResponse.FullName)
	assert.Equal(t, user.EmailAddress, userResponse.EmailAddress)
	assert.Equal(t, user.Timestamp, userResponse.Timestamp)
}

func TestUserRepositoryImpl_Delete(t *testing.T) {
	err := userRepository.Delete(conn, user.ID)
	assert.Nil(t, err)
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	err := SaveOrUpdateUser(user)
	assert.Nil(t, err)

	isEmailExist := userRepository.FindByEmail(conn, user.EmailAddress)
	assert.Equal(t, true, isEmailExist)

	err = userRepository.Delete(conn, user.ID)
	assert.Nil(t, err)
}
