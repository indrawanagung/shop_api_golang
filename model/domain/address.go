package domain

type Address struct {
	ID         string `gorm:"primary_key;column:id"`
	Name       string `gorm:"column:name"`
	CityID     string `gorm:"column:city_id"`
	PostalCode string `gorm:"column:postal_code"`
	UserID     string `gorm:"column:user_id"`
	City       *City  `gorm:"foreignKey:city_id;references:id"`
	User       User   `gorm:"foreignKey:user_id;references:id"`
	Timestamp
}

func (a *Address) TableName() string {
	return "address"
}
