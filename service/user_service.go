package service

import "github.com/indrawanagung/shop_api_golang/model/web"

type UserServiceInterface interface {
	Save(request web.UserCreateOrUpdateRequest) string
	Update(ID string, request web.UserCreateOrUpdateRequest)
	FindByID(ID string) web.UserResponse
	Delete(ID string)
}
