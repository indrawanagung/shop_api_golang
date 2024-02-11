package domain

import "gorm.io/gorm"

type City struct {
	ID        string         `gorm:"primary_key;column:id"`
	City      string         `gorm:"column:city"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (c *City) TableName() string {
	return "cities"
}
