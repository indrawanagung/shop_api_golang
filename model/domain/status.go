package domain

type Status struct {
	ID               string `gorm:"primaryKey;column:id"`
	CategoryStatusID string `gorm:"column:category_status_id"`
	Name             string `gorm:"column:name"`
	Timestamp
	CategoryStatus CategoryStatus `gorm:"foreignKey:category_status_id;references:id"`
}

func (s *Status) TableName() string {
	return "status"
}
