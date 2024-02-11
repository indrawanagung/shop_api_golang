package domain

type WeightUnit struct {
	ID       string `gorm:"primaryKey;column:id"`
	UnitName string `gorm:"unit_name"`
}

func (w *WeightUnit) TableName() string {
	return "weight_units"
}
