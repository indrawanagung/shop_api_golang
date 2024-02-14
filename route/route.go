package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/indrawanagung/shop_api_golang/controller"
	"github.com/indrawanagung/shop_api_golang/exception"
)

func New(
	userController controller.UserControllerInterface,
	cityController controller.CityControllerInterface,
	addressController controller.AddressControllerInterface,

) *fiber.App {

	app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})
	app.Use(recover.New())

	api := app.Group("/shop-api")
	v1 := api.Group("/v1")

	v1.Get("/users/:userID", userController.FindByID)
	v1.Post("/users", userController.Save)
	v1.Put("/users/:userID", userController.Update)
	v1.Delete("/users/:userID", userController.Delete)

	v1.Get("/cities", cityController.FindAll)
	v1.Get("/cities/:cityID", cityController.FindByID)

	v1.Get("/address/:addressID", addressController.FindByID)
	v1.Post("/address", addressController.Save)
	v1.Put("/address/:addressID", addressController.Update)
	v1.Delete("/address/:addressID", addressController.Update)

	// Error Handler

	return app
}
