package domain

type ShoppingCartItems struct {
	ID             string `gorm:"primaryKey"`
	ShoppingCartID string
	ProductID      string
	Qty            int
	Timestamp
	ShoppingCart ShoppingCart `gorm:"foreignKey:shopping_cart_id:references:id"`
	Product      Product      `gorm:"foreignKey:product_id;references:id"`
}

func (i *ShoppingCartItems) TableName() string {
	return "shopping_cart_items"
}
