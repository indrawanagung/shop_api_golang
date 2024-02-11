package domain

type ShopOrderItem struct {
	ID          string `gorm:"primaryKey"`
	ProductID   string
	ShopOrderID string
	Qty         int
	Price       int64
	StatusID    string
	Product     Product   `gorm:"foreignKey:product_id;references:id"`
	ShopOrder   ShopOrder `gorm:"foreignKey:shop_order_id;references:id"`
	Status      Status    `gorm:"foreignKey:status_id;references:id"`
}

func (s *ShopOrderItem) TableName() string {
	return "shop_order_items"
}
