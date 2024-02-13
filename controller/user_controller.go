package controller

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	Register(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
}

type UserControllerInterface interface {
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
