package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/db"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/indrawanagung/shop_api_golang/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func DB() *gorm.DB {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	conn := db.OpenConnection(config.DBSource)
	return conn
}

var database = DB()
var validate = validator.New()
var userRepository = repository.NewUserRepository()
var userService = NewUserService(database, userRepository, validate)

var userRequest web.UserCreateOrUpdateRequest = web.UserCreateOrUpdateRequest{
	FullName:     "Golang",
	EmailAddress: "golang@gmail.com",
	PhoneNumber:  "0891231235432",
	Password:     "123123123",
}

func TestUserServiceImpl_FindByID(t *testing.T) {
	id := userService.Save(userRequest)

	userResponse := userService.FindByID(id)
	assert.NotNil(t, userResponse)

	userService.Delete(id)
}

func TestUserServiceImpl_Update(t *testing.T) {
	id := userService.Save(userRequest)
	userService.Update(id, userRequest)

	user := userService.FindByID(id)
	assert.NotNil(t, user)
	assert.Equal(t, userRequest.EmailAddress, user.EmailAddress)

	userService.Delete(id)
}

func TestUserServiceImpl_Delete(t *testing.T) {
	id := userService.Save(userRequest)
	userService.Delete(id)
}
