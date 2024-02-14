package service

import (
	"fmt"
	"github.com/indrawanagung/shop_api_golang/model/web"
	"github.com/stretchr/testify/assert"
	"testing"
)

var authService = NewAuthService(userRepository, database, validate)

func TestAuthServiceImpl_LoginFailed(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	request := web.LoginRequest{
		Email:    "firefox",
		Password: "123",
	}
	login := authService.Login(request)
	assert.NotNil(t, login)
}

func TestAuthServiceImpl_LoginSuccess(t *testing.T) {
	request := web.LoginRequest{
		Email:    "golang@gmail.com",
		Password: "123123",
	}
	login := authService.Login(request)
	assert.NotNil(t, login)
	fmt.Println(login)
}
