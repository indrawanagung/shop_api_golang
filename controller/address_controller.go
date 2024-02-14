package controller

import "github.com/gofiber/fiber/v2"

type AddressControllerInterface interface {
	FindAllByUserID(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
