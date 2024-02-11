package domain

type Product struct {
	ID                string `gorm:"primaryKey;column:id"`
	Name              string `gorm:"column:name"`
	ProductCategoryID string `gorm:"column:product_category_id"`
	StatusID          string `gorm:"column:status_id"`
	Price             int64  `gorm:"column:price"`
	Timestamp
	ProductCategory ProductCategory `gorm:"foreignKey:product_category_id;references:id"`
	Status          Status          `gorm:"foreignKey:status_id;references:id"`
}

func (p *Product) TableName() string {
	return "products"
}
