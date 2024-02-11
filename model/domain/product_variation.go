package domain

type ProductVariation struct {
	ProductID         string          `gorm:"primaryKey"`
	VariationOptionID string          `gorm:"primaryKey"`
	Product           Product         `gorm:"foreignKey:product_id;references:id"`
	VariationOption   VariationOption `gorm:"foreignKey:variation_option_id;references:id"`
}

func (p *ProductVariation) TableName() string {
	return "product_variations"
}
