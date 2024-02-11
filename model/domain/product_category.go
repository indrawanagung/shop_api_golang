package domain

type ProductCategory struct {
	ID           string `gorm:"primaryKey;column:id"`
	CategoryName string `gorm:"column:category_name"`
}

func (p *ProductCategory) name() string {
	return "product_categories"
}
