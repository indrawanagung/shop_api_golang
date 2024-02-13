package controller

import "github.com/gofiber/fiber/v2"

type CityControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
}
