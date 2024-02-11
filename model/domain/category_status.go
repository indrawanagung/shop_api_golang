package domain

type CategoryStatus struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
}

func (c *CategoryStatus) TableName() string {
	return "category_status"
}
