package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/indrawanagung/shop_api_golang/model/domain"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/indrawanagung/shop_api_golang/util"
	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	Database          *gorm.DB
	AddressRepository repository.AddressRepositoryInterface
	Validate          *validator.Validate
}

func NewAddressService(database *gorm.DB, addressRepository repository.AddressRepositoryInterface, validate *validator.Validate) AddressServiceInterface {
	return &AddressServiceImpl{
		Database:          database,
		AddressRepository: addressRepository,
		Validate:          validate,
	}
}

func (s AddressServiceImpl) FindAllByUserID(userID string) []web.AddressResponse {
	addresses := s.AddressRepository.FindAllByUserID(s.Database, userID)
	return web.ToAddressResponses(addresses)
}

func (s AddressServiceImpl) FindByID(addressID string) web.AddressResponse {
	err, address := s.AddressRepository.FindByID(s.Database, addressID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToAddressResponse(address)
}

func (s AddressServiceImpl) Save(addressRequest web.AddressCreateOrUpdateRequest) string {
	err := s.Validate.Struct(addressRequest)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	address := domain.Address{
		ID:         util.GenerateUUID(),
		Name:       addressRequest.Name,
		CityID:     addressRequest.CityID,
		PostalCode: addressRequest.PostalCode,
		UserID:     addressRequest.UserID,
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	}

	s.AddressRepository.SaveOrUpdate(s.Database, address)
	return address.ID
}

func (s AddressServiceImpl) Update(addressRequest web.AddressCreateOrUpdateRequest, addressID string) {
	err, d := s.AddressRepository.FindByID(s.Database, addressID)
	if err != nil {
		log.Info(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	address := domain.Address{
		ID:         addressID,
		Name:       addressRequest.Name,
		CityID:     addressRequest.CityID,
		PostalCode: addressRequest.PostalCode,
		UserID:     addressRequest.UserID,
		Timestamp: domain.Timestamp{
			CreatedAt: d.CreatedAt,
			UpdatedAt: util.GetUnixTimestamp(),
		},
	}

	s.AddressRepository.SaveOrUpdate(s.Database, address)

}

func (s AddressServiceImpl) Delete(addressID string) {
	err, _ := s.AddressRepository.FindByID(s.Database, addressID)
	if err != nil {
		log.Info(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	s.AddressRepository.Delete(s.Database, addressID)
}
