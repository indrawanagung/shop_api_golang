package route

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/indrawanagung/shop_api_golang/controller"
	"github.com/indrawanagung/shop_api_golang/exception"
	"github.com/spf13/viper"
)

func New(
	userController controller.UserControllerInterface,
	cityController controller.CityControllerInterface,
	addressController controller.AddressControllerInterface,
	authController controller.AuthControllerInterface,

) *fiber.App {
	authentication := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("SECRET_KEY"))},
	})
	app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})
	app.Use(recover.New())
	api := app.Group("/shop-api")
	v1 := api.Group("/v1")

	v1.Post("/auth/login", authController.Login)

	v1.Get("/users/:userID", userController.FindByID)
	v1.Post("/users", userController.Save)
	v1.Put("/users/:userID", userController.Update)
	v1.Delete("/users/:userID", userController.Delete)

	v1.Get("/cities", cityController.FindAll)
	v1.Get("/cities/:cityID", cityController.FindByID)

	v1.Get("/address", authentication, addressController.FindAllByUserID)
	v1.Get("/address/:addressID", authentication, addressController.FindByID)
	v1.Post("/address", addressController.Save)
	v1.Put("/address/:addressID", addressController.Update)
	v1.Delete("/address/:addressID", addressController.Delete)

	// Error Handler

	return app
}
