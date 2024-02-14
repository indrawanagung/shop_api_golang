package web

type AddressCreateOrUpdateRequest struct {
	Name       string `validate:"required,min=1,max=50" json:"name"`
	CityID     string `validate:"required,min=1,max=50" json:"city_id"`
	PostalCode string `validate:"required,min=1,max=50" json:"postal_code"`
	UserID     string `validate:"required,min=1,max=50" json:"user_id"`
}
