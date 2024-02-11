package domain

import "gorm.io/gorm"

type ShopOrder struct {
	ID              string `gorm:"primaryKey"`
	UserID          string
	AddressID       string
	PaymentMethodID string
	TotalPrice      int64
	OrderedAt       string
	DeletedAt       gorm.DeletedAt
	user            User          `gorm:"foreignKey:user_id;references:id"`
	Address         Address       `gorm:"foreignKey:address_id;references:id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:payment_method_id;references:id"`
}

func (s *ShopOrder) TableName() string {
	return "shop_orders"
}
