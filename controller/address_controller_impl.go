package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/service"
	"github.com/indrawanagung/shop_api_golang/util"
)

type AddressControllerImpl struct {
	AddressService service.AddressServiceInterface
}

func NewAddressController(addressService service.AddressServiceInterface) AddressControllerInterface {
	return &AddressControllerImpl{AddressService: addressService}
}

func (c AddressControllerImpl) FindAllByUserID(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	addressesResponse := c.AddressService.FindAllByUserID(userID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   addressesResponse,
	}

	return ctx.Status(200).JSON(webResponse)
}

func (c AddressControllerImpl) FindByID(ctx *fiber.Ctx) error {
	addressID := ctx.Params("addressID")
	addressResponse := c.AddressService.FindByID(addressID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   addressResponse,
	}

	return ctx.Status(200).JSON(webResponse)
}

func (c AddressControllerImpl) Save(ctx *fiber.Ctx) error {
	addressCreateRequest := new(web.AddressCreateOrUpdateRequest)
	if err := ctx.BodyParser(addressCreateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	addressID := c.AddressService.Save(*addressCreateRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"address_id": addressID},
	}

	return ctx.Status(201).JSON(webResponse)
}

func (c AddressControllerImpl) Update(ctx *fiber.Ctx) error {
	addressID := ctx.Params("addressID")
	addressUpdateRequest := new(web.AddressCreateOrUpdateRequest)
	if err := ctx.BodyParser(addressUpdateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	c.AddressService.Update(*addressUpdateRequest, addressID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}

	return ctx.Status(200).JSON(webResponse)
}

func (c AddressControllerImpl) Delete(ctx *fiber.Ctx) error {
	addressID := ctx.Params("addressID")
	c.AddressService.Delete(addressID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}
