package domain

type Store struct {
	ID     string `gorm:"primaryKey;column:id"`
	UserID string `gorm:"column:user_id"`
	Name   string `gorm:"column:name"`
	Points int64  `gorm:"column:points"`
	Timestamp
	User User `gorm:"foreignKey:user_id;references:id"`
}

func (s *Store) TableName() string {
	return "stores"
}
