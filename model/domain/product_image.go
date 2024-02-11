package domain

type ProductImage struct {
	ID        string `gorm:"primaryKey;column:id"`
	ProductID string `gorm:"column:product_id"`
	ImageURL  string `gorm:"column:image_url"`
}

func (p *ProductImage) TableName() string {
	return "product_images"
}
