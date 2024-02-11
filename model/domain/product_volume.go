package domain

type ProductVolume struct {
	ID        string `gorm:"primaryKey;column:id"`
	ProductID string `gorm:"column:product_id"`
	Width     int    `gorm:"column:width"`
	Height    int    `gorm:"column:height"`
	Length    int    `gorm:"column:length"`
	Timestamp
	Product Product `gorm:"foreignKey:product_id;references:id"`
}

func (p *ProductVolume) TableName() string {
	return "product_volumes"
}
