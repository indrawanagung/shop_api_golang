package service

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/spf13/viper"
	"time"
)

type AuthServiceImpl struct {
}

func NewAuthService() AuthServiceInterface {
	return &AuthServiceImpl{}
}
func (a AuthServiceImpl) Login(request web.LoginRequest) interface{} {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":  request.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return t
}
