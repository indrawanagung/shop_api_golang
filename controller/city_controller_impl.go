package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/service"
	"github.com/indrawanagung/shop_api_golang/util"
)

type CityControllerImpl struct {
	CityService service.CityServiceInterface
}

func NewCityController(cityService service.CityServiceInterface) CityControllerInterface {
	return &CityControllerImpl{CityService: cityService}
}

func (c CityControllerImpl) FindAll(ctx *fiber.Ctx) error {
	cities := c.CityService.FindAll()

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   cities,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CityControllerImpl) FindByID(ctx *fiber.Ctx) error {
	cityID := ctx.Params("cityID")
	city := c.CityService.FindByID(cityID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   city,
	}
	return ctx.Status(200).JSON(webResponse)
}
