package domain

type Warehouses struct {
	ID        string `gorm:"primaryKey;column:id"`
	StoreID   string `gorm:"column:store_id"`
	AddressID string `gorm:"column:address_id"`
	StatusID  string `gorm:"status_id"`
	Name      string `gorm:"name"`
	Timestamp
	Store   Store   `gorm:"foreignKey:store_id;references:id"`
	Address Address `gorm:"foreignKey:address_id;references:id"`
	Status  Status  `gorm:"foreignKey:status_id;references:id"`
}

func (w *Warehouses) TableName() string {
	return "warehouses"
}
