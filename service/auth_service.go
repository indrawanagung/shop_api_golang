package service

import "github.com/indrawanagung/shop_api_golang/model/web"

type AuthServiceInterface interface {
	Login(request web.LoginRequest) interface{}
}
