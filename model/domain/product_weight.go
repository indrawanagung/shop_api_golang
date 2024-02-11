package domain

type ProductWeight struct {
	ID           string `gorm:"primaryKey;column:id"`
	ProductID    string `gorm:"column:product_id"`
	WeightUnitID string `gorm:"column:weight_unit_id"`
	Value        int    `gorm:"column:value"`
	Timestamp
	Product    Product    `gorm:"foreignKey:product_id;references:id"`
	WeightUnit WeightUnit `gorm:"foreignKey:weight_unit_id;references:id"`
}

func (p *ProductWeight) TableName() string {
	return "product_weights"
}
