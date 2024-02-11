package domain

type PaymentType struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

func (p *PaymentType) TableName() string {
	return "payment_types"
}
