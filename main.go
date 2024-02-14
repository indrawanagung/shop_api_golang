package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/indrawanagung/shop_api_golang/controller"
	"github.com/indrawanagung/shop_api_golang/db"
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/indrawanagung/shop_api_golang/route"
	"github.com/indrawanagung/shop_api_golang/service"
	"github.com/indrawanagung/shop_api_golang/util"
	"os"
)

func main() {

	validate := validator.New()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	database := db.OpenConnection(config.DBSource)

	userRepository := repository.NewUserRepository()
	cityRepository := repository.NewCityRepository()
	addressRepository := repository.NewAddressRepository()

	userService := service.NewUserService(database, userRepository, validate)
	cityService := service.NewCityService(cityRepository, database)
	addressService := service.NewAddressService(database, addressRepository, validate)
	authService := service.NewAuthService(userRepository, database, validate)

	userController := controller.NewUserController(userService)
	cityController := controller.NewCityController(cityService)
	addressController := controller.NewAddressController(addressService)
	authController := controller.NewAuthController(authService)

	app := route.New(userController, cityController, addressController, authController)
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	f, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(f)
	log.Info("server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
