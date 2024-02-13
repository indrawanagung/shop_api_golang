package web

import "github.com/indrawanagung/shop_api_golang/model/domain"

type UserCreateOrUpdateRequest struct {
	FullName     string `validate:"required,min=5,max=50" json:"full_name"`
	EmailAddress string `validate:"required,min=5,max=50" json:"email_address"`
	PhoneNumber  string `validate:"required,min=7,max=50" json:"phone_number"`
	Password     string `validate:"required,min=5,max=50" json:"password"`
}

type UserResponse struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email-address"`
	PhoneNumber  string `json:"phone_number"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress,
		PhoneNumber:  user.PhoneNumber,
	}
}
