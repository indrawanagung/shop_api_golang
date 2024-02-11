package domain

type Variation struct {
	ID                string `gorm:"primaryKey;column:id"`
	ProductCategoryID string `gorm:"column:product_category_id"`
	VariationName     string `gorm:"column:variation_name"`
	Timestamp
	ProductCategory ProductCategory `gorm:"foreignKey:product_category_id;references:id"`
}

func (v *Variation) TableName() string {
	return "variations"
}
