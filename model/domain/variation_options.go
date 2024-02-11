package domain

type VariationOption struct {
	ID          string `gorm:"primaryKey"`
	VariationID string
	OptionName  string
	Description string
	Timestamp
	Variation Variation `gorm:"foreignKey:variation_id;references:id"`
}

func (v *VariationOption) TableName() string {
	return "variation_options"
}
