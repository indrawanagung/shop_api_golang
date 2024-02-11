package domain

type ProductStock struct {
	ID          string `gorm:"primaryKey;column:id"`
	ProductID   string `gorm:"column:product_id"`
	Total       int    `gorm:"column:total"`
	NotReserved int    `gorm:"not_reserved"`
	Reserved    int    `gorm:"reserved"`
	Timestamp
	Product Product `gorm:"foreignKey:product_id;references:id"`
}

func (p *ProductStock) TableName() string {
	return "product_stocks"
}
