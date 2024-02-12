package web

type LoginRequest struct {
	Email    string `validate:"required,min=1,max=50" json:"email"`
	Password string `validate:"required,min=1,max=50" json:"password"`
}

type LoginResponse struct {
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
	Token        string `json:"token"`
}
