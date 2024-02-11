package domain

type ProductView struct {
	ProductID string  `gorm:"primaryKey;column:product_id"`
	UserID    string  `gorm:"primaryKey;column:user_id"`
	Product   Product `gorm:"foreignKey:product_id;references:id"`
	User      User    `gorm:"foreignKey:user_id;references:id"`
}

func (p *ProductView) TableName() string {
	return "product_views"
}
