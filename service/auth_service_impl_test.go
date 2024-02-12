package service

import (
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthServiceImpl_Login(t *testing.T) {
	authService := NewAuthService()

	request := web.LoginRequest{
		Email:    "firefox",
		Password: "123",
	}
	login := authService.Login(request)
	assert.NotNil(t, login)
}
