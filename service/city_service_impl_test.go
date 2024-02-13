package service

import (
	"github.com/indrawanagung/shop_api_golang/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cityRepository = repository.NewCityRepository()
var cityService = NewCityService(cityRepository, database)

func TestCityServiceImpl_FindAll(t *testing.T) {
	cities := cityService.FindAll()
	assert.NotEqual(t, 0, len(cities))
}

func TestCityServiceImpl_FindByID(t *testing.T) {
	cityID := "1"
	city := cityService.FindByID(cityID)
	assert.NotNil(t, city)
	assert.Equal(t, "makassar", city.City)
}
