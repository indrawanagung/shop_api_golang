package domain

type User struct {
	ID           string `gorm:"primary_key;column:id"`
	FullName     string `gorm:"column:fullname"`
	EmailAddress string `gorm:"column:email_address"`
	PhoneNumber  string `gorm:"column:phone_number"`
	Password     string `gorm:"column:password"`
	Timestamp
	Addresses []Address `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
