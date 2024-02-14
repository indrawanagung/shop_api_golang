package web

import "github.com/indrawanagung/shop_api_golang/model/domain"

type AddressCreateOrUpdateRequest struct {
	Name       string `validate:"required,min=1,max=50" json:"name"`
	CityID     string `validate:"required,min=1,max=50" json:"city_id"`
	PostalCode string `validate:"required,min=1,max=50" json:"postal_code"`
	UserID     string `validate:"required,min=1,max=50" json:"user_id"`
}

type AddressResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func ToAddressResponse(address domain.Address) AddressResponse {
	return AddressResponse{
		ID:         address.ID,
		Name:       address.Name,
		City:       address.City.City,
		PostalCode: address.PostalCode,
	}
}

func ToAddressResponses(addresses []domain.Address) []AddressResponse {
	var adddressResponses []AddressResponse

	for _, address := range addresses {
		adddressResponses = append(adddressResponses, ToAddressResponse(address))
	}
	return adddressResponses
}
