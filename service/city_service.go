package service

import "github.com/indrawanagung/shop_api_golang/model/domain"

type CityServiceInterface interface {
	FindAll() []domain.City
	FindByID(cityID string) domain.City
}
