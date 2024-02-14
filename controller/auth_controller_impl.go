package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/service"
	"github.com/indrawanagung/shop_api_golang/util"
)

type AuthControllerImpl struct {
	AuthService service.AuthServiceInterface
}

func NewAuthController(authService service.AuthServiceInterface) AuthControllerInterface {
	return &AuthControllerImpl{AuthService: authService}
}

func (c AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	loginRequest := new(web.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	loginResponse := c.AuthService.Login(*loginRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   loginResponse,
	}
	return ctx.Status(200).JSON(webResponse)

}
