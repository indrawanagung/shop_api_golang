package domain

type ShoppingCart struct {
	ID     string `gorm:"primaryKey"`
	UserID string
	User   User `gorm:"foreignKey:user_id;references:id"`
}

func (s *ShoppingCart) TableName() string {
	return "shopping_carts"
}
