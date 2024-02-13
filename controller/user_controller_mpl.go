package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/indrawanagung/shop_api_golang/service"
	"github.com/indrawanagung/shop_api_golang/util"
)

type UserControllerImpl struct {
	UserService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u UserControllerImpl) Save(c *fiber.Ctx) error {
	userCreateRequest := new(web.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userCreateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userID := u.UserService.Save(*userCreateRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"user_id": userID},
	}
	return c.Status(201).JSON(webResponse)
}

func (u UserControllerImpl) Update(c *fiber.Ctx) error {
	userID := c.Params("userID")
	userUpdateRequest := new(web.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userUpdateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	u.UserService.Update(userID, *userUpdateRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return c.Status(200).JSON(webResponse)
}

func (u UserControllerImpl) FindByID(c *fiber.Ctx) error {
	userID := c.Params("userID")

	userResponse := u.UserService.FindByID(userID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   userResponse,
	}

	return c.Status(200).JSON(webResponse)
}

func (u UserControllerImpl) Delete(c *fiber.Ctx) error {
	userID := c.Params("userID")

	u.UserService.Delete(userID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}

	return c.Status(200).JSON(webResponse)
}
