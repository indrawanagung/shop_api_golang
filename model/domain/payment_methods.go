package domain

type PaymentMethod struct {
	ID            string `gorm:"primaryKey"`
	PaymentTypeID string
	AccountNumber string
	ExpiryDate    string
	StatusID      string
	Timestamp
	PaymentType PaymentType `gorm:"foreignKey:payment_type_id;references:id"`
	Status      Status      `gorm:"foreignKey:status_id;references:id"`
}

func (p *PaymentMethod) TableName() string {
	return "payment_methods"
}
