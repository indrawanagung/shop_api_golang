package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cityRepository = NewCityRepository()

func TestCityRepositoryImpl_FindAll(t *testing.T) {
	err, cities := cityRepository.FindAll(conn)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(cities))
}

func TestCityRepositoryImpl_FindByID(t *testing.T) {
	cityID := "1"
	err, city := cityRepository.FindByID(conn, cityID)
	assert.Nil(t, err)
	assert.NotNil(t, city)
	assert.Equal(t, "makassar", city.City)
}
