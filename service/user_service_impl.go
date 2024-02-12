package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/indrawanagung/shop_api_golang/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Database       *gorm.DB
	UserRepository repository.UserRepositoryInterface
	Validate       *validator.Validate
}

func NewUserService(database *gorm.DB, userRepository repository.UserRepositoryInterface, validate *validator.Validate) UserServiceInterface {
	return &UserServiceImpl{
		Database:       database,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s UserServiceImpl) Save(request web.UserCreateOrUpdateRequest) string {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	id := util.GenerateUUID()
	err = s.UserRepository.SaveOrUpdate(s.Database, domain.User{
		ID:           id,
		FullName:     request.FullName,
		EmailAddress: request.EmailAddress,
		PhoneNumber:  request.PhoneNumber,
		Password:     string(passwordHash),
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	return id
}

func (s UserServiceImpl) Update(ID string, request web.UserCreateOrUpdateRequest) {

	err, user := s.UserRepository.FindByID(s.Database, ID)
	if err != nil {
		log.Fatal(err)
	}

	var passwordHash string
	if request.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		passwordHash = string(hash)
	} else {
		passwordHash = user.Password
		request.Password = util.GetUnixTimestamp() // random string, just for pass validate password request
	}

	err = s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	err = s.UserRepository.SaveOrUpdate(s.Database, domain.User{
		ID:           ID,
		FullName:     request.FullName,
		EmailAddress: request.EmailAddress,
		PhoneNumber:  request.PhoneNumber,
		Password:     passwordHash,
		Timestamp: domain.Timestamp{
			UpdatedAt: util.GetUnixTimestamp(),
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (s UserServiceImpl) FindByID(ID string) web.UserResponse {
	err, user := s.UserRepository.FindByID(s.Database, ID)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("user id %s is not found", ID)))
	}

	return web.ToUserResponse(user)
}

func (s UserServiceImpl) Delete(ID string) {
	err, _ := s.UserRepository.FindByID(s.Database, ID)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("user id %s is not found", ID)))
	}

	err = s.UserRepository.Delete(s.Database, ID)
	if err != nil {
		log.Fatal(err)
	}
}